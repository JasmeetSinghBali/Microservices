package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jasmeet-1998/Microservices/structuring_microservice_basics/handlers"
)

func main() {
	/*
		new logger gl instance
		with prefix Orders-api and standard flags(LstdFlags)
	*/
	gl := log.New(os.Stdout, "Orders-api", log.LstdFlags)

	/*
		reff to the greetings handler gh
		injecting the logger into the NewGreetings method implementing the traces/logger instance of Greetings struct for greeting interface
	*/
	gh := handlers.NewGreetings(gl)
	bh := handlers.NewDasvadania(gl)

	/*
		registering the greetings handler with the servemux with servermux sm instance
		for pattern /
	*/
	sm := http.NewServeMux()
	sm.Handle("/", gh)
	sm.Handle("/bye", bh)

	/*passsing the servemux instance to the server*/
	http.ListenAndServe("127.0.0.1:5000", sm)
}
