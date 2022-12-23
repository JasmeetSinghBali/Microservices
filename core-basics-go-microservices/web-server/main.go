package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*entry point for any golang package*/
func main() {

	/*simplistic handler for pattern '/' that logs Hello World in terminal*/
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("ssa")

		/*
			since the request Body implements ioreader , ioutil can be used to read the entire request body
			data now holds the entire body
		*/
		data, err := ioutil.ReadAll(r.Body)

		/* handle error naive way*/
		// if err != nil {
		// 	/*set httpStatusCode with WriteHeader*/
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// 	/*return data as slice of byte data type*/
		// 	rw.Write([]byte("Something Went Wrong!!"))
		// 	return
		// }

		/*handle error a little less naive*/
		// ref: https://pkg.go.dev/net/http@go1.19.4#Error
		if err != nil {
			http.Error(rw, "Something Went Wrong!", http.StatusBadRequest)
			return
		}

		/*
			Print data as string in terminal where go code compiled
			$ curl -v -d 'Jas' localhost:5000
			Data Request Body: Jas
		*/
		log.Printf("Data Request Body: %s", data)

		/*
			print backs to user via iowriter(fmt.Fprintf) response writer interface where curl request was made
			$ curl -v -d 'Jas' localhost:5000
		*/
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("dasvadania")
	})
	/*
	  web server listening on localhost:5000 for allIP with port 5000 mention ':5000'
	  nil is replaced here by default servemux when code is compiled ref: https://go.dev/src/net/http/server.go?s=61509%3A61556#L2378
	*/
	http.ListenAndServe("127.0.0.1:5000", nil)
}
