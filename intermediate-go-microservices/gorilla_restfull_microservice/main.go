package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/handlers"
	"github.com/gorilla/mux"
)

func main() {

	gl := log.New(os.Stdout, "Product-api", log.LstdFlags)

	pl := handlers.NewProducts(gl)

	// create a serve mux via gorilla web toolkit (Package- gorilla/mux)
	sm := mux.NewRouter()

	// 📝 ref: https://pkg.go.dev/github.com/gorilla/mux#section-readme subrouter section
	// creates a sub-router named getRouter instance that has attached only GET type routes with it
	getRouter := sm.Methods(http.MethodGet).Subrouter()

	// 📝 now GET routes can be attached to this getRouter sub-router instance
	getRouter.HandleFunc("/products", pl.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", pl.GetProduct)

	// put sub-router instance named putRouter that has PUT type http request associated & mapped to it only
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	// 📝 the route attached has id as param which is regex 0-9 and can be 1 or more & gets auto extracted by gorilla router
	putRouter.HandleFunc("/{id:[0-9]+}", pl.UpdateProducts)
	// 📝 adding validation middleware to putRouter subrouter that gets executed before the handler
	putRouter.Use(pl.ProductValidationMiddleware)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", pl.AddProduct)
	// 📝 adding validation middleware to postRouter subrouter that gets executed before the handler
	postRouter.Use(pl.ProductValidationMiddleware)

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
