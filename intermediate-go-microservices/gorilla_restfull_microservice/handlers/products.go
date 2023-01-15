package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
	"github.com/gorilla/mux"
)

// Key: value based context http struct
type KeyProduct struct{}

// products struct for fetching/updating products with logger,validator as tracer & checker
type Products struct {
	tracer  *log.Logger
	checker *data.Validation
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// generic error returned by the server
type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

// NewProducts handler returns new Products struct instance product with tracer & checker
func NewProducts(tracer *log.Logger, checker *data.Validation) *Products {
	return &Products{tracer, checker}
}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
