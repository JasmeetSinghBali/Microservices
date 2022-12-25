package handlers

import (
	"log"
	"net/http"
)

type Dasvadania struct {
	traces *log.Logger
}

func NewDasvadania(traces *log.Logger) *Dasvadania {

	return &Dasvadania{traces}

}

func (d *Dasvadania) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("Until next time"))

}
