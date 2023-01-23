package handlers

import (
	"log"
	"net/http"
	"path/filepath"

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

// method- ServeHTTP to Files struct, to implement http.Handler interface
func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.tracer.Println("Handle POST id|filename", id, fn)

	//check filepath validity for name & file
	if id == "" || fn == "" {
		f.invalidURI(r.URL.String(), rw)
		return
	}

	f.saveFile(id, fn, rw, r)
}

// method- invalidURI to Files struct, return error response if invalid uri was provided
func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.tracer.Fatal("Invalid path", uri)
	http.Error(rw, "Invalid file path should be in format: /[id]/[filepath]", http.StatusBadRequest)
}

//method - saveFile to Files struct
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.tracer.Println("Save the file id|path", id, path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.tracer.Fatal("failed to store file", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
