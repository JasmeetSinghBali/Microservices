package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Greetings struct with traces i.e logger object that can be replaced by any object while testing,
sort of DEPENDENCY INJECTION STRUCTURING
*/
type Greetings struct {
	traces *log.Logger
}

func NewGreetings(traces *log.Logger) *Greetings {
	return &Greetings{traces}
}

/*
method of Greetings struct that implements the http handler interface
*/
func (g *Greetings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	/* a common logger dependency injection instance Greetings can be used i.e traces in the Greetings function handler method*/
	g.traces.Println("ssa")

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something Went Wrong!", http.StatusBadRequest)
		return
	}

	log.Printf("Data Request Body: %s", data)

	fmt.Fprintf(rw, "Hello %s", data)
}
