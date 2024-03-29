package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	gl := log.New(os.Stdout, "Product-api", log.LstdFlags)
	gv := data.NewValidation()

	// instantiate handler wid new logger & validator instance respectively
	pl := handlers.NewProducts(gl, gv)

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
	putRouter.HandleFunc("/products", pl.UpdateProduct)
	// 📝 adding validation middleware to putRouter subrouter that gets executed before the handler
	putRouter.Use(pl.ProductValidationMiddleware)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", pl.AddProduct)
	// 📝 adding validation middleware to postRouter subrouter that gets executed before the handler
	postRouter.Use(pl.ProductValidationMiddleware)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", pl.Delete)

	// ✨ redoc is used to serve the swagger handler ref: https://github.com/Redocly/redoc
	redocOptions := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerDocMiddleware := middleware.Redoc(redocOptions, nil)
	getRouter.Handle("/docs", swaggerDocMiddleware)
	// ✨ add specific handler to serve swagger.yaml to load and serve it when /docs is hit
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// 📝 CORS, cors handler ch
	// allowing localhost:3000 where frontend (in react) is running on
	// 💡 for public api mention []string["*"] wild card to allow request from anywhere
	// ✨ if cookies is used for authentication then access-control-allow-credentials:true should be added
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	s := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      ch(sm), // wrap the entire serve mux with the cors handler to allow only 3000 to make request to go api at 5000
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
