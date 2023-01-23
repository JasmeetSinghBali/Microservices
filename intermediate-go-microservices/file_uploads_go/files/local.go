package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local- implementation of storage interface defined in storage.go
// for local disk storage
type Local struct {
	maxFileSize int
	basePath    string
}

// NewLocal- method to Local struct creates a new local filesystem with basePath provided
//basePath- file dir to store files, maxSize allowed size i.e number of bytes
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil

}

// returns the absolute  full path to the relative path provided as path
func (l *Local) fullPath(path string) string {
	//append given reltive path with base path
	return filepath.Join(l.basePath, path)
}

// Save- method to Local struct, to save contents to the writer where path is relative path
func (l *Local) Save(path string, contents io.Reader) error {
	// ull path for the file
	fp := l.fullPath(path)

	//get dir and check its existance
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}

	// if file alredy exist delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("Unable to delte already existing file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// other than not exist error
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	//create new file at path
	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer f.Close()

	// write contents to new file, with max size limit check can be done via buffer reader comparing with loop overed content length manually NOTE- content-length header must never be trusted
	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write contents to the file: %w", err)
	}

	return nil

}

// method Get to Local struct
// get the file via given path and return a reader
// the invoking function of this method is responsible to close the reader
func (l *Local) Get(path string) (*os.File, error) {
	//get the full path for the file
	fp := l.fullPath(path)

	//open the file
	f, err := os.Open(fp)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open file: %w", err)
	}

	return f, nil
}
