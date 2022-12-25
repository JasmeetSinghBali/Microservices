package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Greetings struct {
	traces *log.Logger
}

func NewGreetings(traces *log.Logger) *Greetings {
	return &Greetings{traces}
}

func (g *Greetings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	g.traces.Println("ssa")

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something Went Wrong!", http.StatusBadRequest)
		return
	}

	log.Printf("Data Request Body: %s", data)

	fmt.Fprintf(rw, "Hello %s", data)
}
