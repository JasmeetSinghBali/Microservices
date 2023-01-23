package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/file_uploads_go/files"
	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/file_uploads_go/handlers"
	"github.com/gorilla/mux"
)

func main() {

	gl := log.New(os.Stdout, "Product-api", log.LstdFlags)

	// create the storage class, use local storage
	// max filesize 5MB
	stor, err := files.NewLocal("./imagestore", 1024*1000*5)
	if err != nil {
		gl.Panic("Unable to create storage", "error", err)
		os.Exit(1)
	}

	// create the handlers
	fh := handlers.NewFiles(stor, gl)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	// get files
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir("./imagestore"))),
	)

	// create a new server
	s := http.Server{
		Addr:         "127.0.0.1:5000",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     gl,                // the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		gl.Println("Starting server", "bind_address", "./imagestore")

		err := s.ListenAndServe()
		if err != nil {
			gl.Panic("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	gl.Println("Shutting down server with", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
