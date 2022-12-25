package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/product-api/handlers"
)

func main() {

	gl := log.New(os.Stdout, "Product-api", log.LstdFlags)

	gh := handlers.NewGreetings(gl)
	bh := handlers.NewDasvadania(gl)
	pl := handlers.NewProducts(gl)

	sm := http.NewServeMux()
	sm.Handle("/", gh)
	sm.Handle("/bye", bh)
	sm.Handle("/products", pl)

	s := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      sm,
		ErrorLog:     gl,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		gl.Println("Started server on port 5000")
		err := s.ListenAndServe()
		if err != nil {
			gl.Printf("Error starting server %s\n", err)
		}
		os.Exit(1)
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	gl.Println("Recieved a interupt/kill signal from sigChan (channel) , gracefully shutting down", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)
}
