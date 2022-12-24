package main

import (
	"log"
	"net/http"
	"os"
	"time"

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

	/*
		Tuning the server
		reff: https://pkg.go.dev/net/http#Server
	*/
	s := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
