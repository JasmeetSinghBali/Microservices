package handlers

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/file_uploads_go/files"
	"github.com/gorilla/mux"
)

/*Files - handler for reading & writing files*/
type Files struct {
	tracer *log.Logger
	store  files.Storage
}

// create new file handler
func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{store: s, tracer: l}
}

// try curl localhost:5000/files/1/test.png --data-binary @test.png
// method- UploadREST to Files struct, to implement http.Handler interface
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.tracer.Println("Handle POST id|filename", id, fn)

	//check filepath validity for name & file
	if id == "" || fn == "" {
		f.invalidURI(r.URL.String(), rw)
		return
	}

	f.saveFile(id, fn, rw, r.Body)
}

// method- UploadMultiPart to Files struct, to implement mutipart uploads
// 💡 refF: https://pkg.go.dev/net/http#Request.ParseForm
func (f *Files) UploadMultiPart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.tracer.Panic("Bad request", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	// ref: https://pkg.go.dev/net/http#Request.FormFile
	// <form action="http://localhost:3000" method="post" enctype="multipart/form-data">
	// <input type="text" name="id" value="">
	// <input type="file" name="file">
	// <button type="submit">Submit</button>
	// </form>
	id, idErr := strconv.Atoi(r.FormValue("id"))
	f.tracer.Println("Process form for id", id)
	// validation & converstion of the grabbed id value from client
	if idErr != nil {
		f.tracer.Panic("Bad request", idErr)
		http.Error(rw, "Expected integer id", http.StatusBadRequest)
		return
	}

	// ff is file and mh is header attached to it that has meta data for file
	ff, mh, err := r.FormFile("file")
	if err != nil {
		f.tracer.Panic("Bad request", err)
		http.Error(rw, "Expected file", http.StatusBadRequest)
		return
	}

	// save file
	f.saveFile(r.FormValue("id"), mh.Filename, rw, ff)
}

// method- invalidURI to Files struct, return error response if invalid uri was provided
func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.tracer.Fatal("Invalid path", uri)
	http.Error(rw, "Invalid file path should be in format: /[id]/[filepath]", http.StatusBadRequest)
}

//method - saveFile to Files struct
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.tracer.Println("Save the file id|path", id, path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.tracer.Fatal("failed to store file", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
