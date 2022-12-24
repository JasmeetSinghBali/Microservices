package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	/*
		wrapped serve mux with custom tuning in goroutine
		so it does not block the main flow of the program
	*/
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			gl.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// passing value from channel sigChan to a variable via <- (kinda dequeing operator to dequeue the value from the channel)
	sig := <-sigChan
	gl.Println("Recieved a interupt/kill signal from sigChan (channel) , gracefully shutting down", sig)

	/*
		wait for the request currently been processed,
		and from this point onwards i.e calling Shutdown() ,
		server wont take any more request and after current req are processed
		it shuts down the server
	*/
	// timeout context for server shutdown
	// so their will be 30 second time-buffer for server to process exisiting request, if
	// after 30 seconds request still pending then forcefully close it
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)
}
