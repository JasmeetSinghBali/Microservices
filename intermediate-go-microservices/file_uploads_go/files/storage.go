package files

import "io"

//interface method to save file , implementation can be of type local or cloud storage
type Storage interface {
	Save(path string, file io.Reader) error
}
