package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
	"github.com/gorilla/mux"
)

// Key: value based context http struct
type KeyProduct struct{}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to be deleted from db
	// in:path
	// required:true
	ID int `json:"id"`
}

type Products struct {
	tracer *log.Logger
}

// generic error returned by the server
type GenericError struct {
	Message string `json:"message"`
}

func NewProducts(tracer *log.Logger) *Products {
	return &Products{tracer}
}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}

/*
Middleware that validates the payload while creating or updating product
*/
func (p Products) ProductValidationMiddleware(next http.Handler) http.Handler {
	// response is just an interface, where request is reff to the actual request passed on from client
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prd := data.Product{}

		err := prd.FromJSON(r.Body)
		if err != nil {
			p.tracer.Println("Failed to parse json & go oriented data/struct", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// 📝 validate/sanitize the product paylaod via Validate method exposed via data package reff: data/products.go
		err = prd.Validate()
		if err != nil {
			p.tracer.Println("Failed to validate/sanitize the payload", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// attaches the product i.e prd the unmarshalled json extracted from request body as newContext on the request object that is common and has reff from client--->middlewares--->handlers--->data access methods
		ctx := context.WithValue(r.Context(), KeyProduct{}, prd)
		// create new req with old common r (req) object reff with attached newContext
		req := r.WithContext(ctx)

		// pass control the http handler with mutated req & interface only rw
		next.ServeHTTP(rw, req)

	})
}
