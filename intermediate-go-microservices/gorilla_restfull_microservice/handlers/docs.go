// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

//
// These types are defined for swagger documentation & have no real significance or use in handlers

// Generic error message returned as string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Description of the validation error
	// in: body
	Body ValidationError
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// single product returned in the response
// swagger:response productResponse
type productResponse struct {
	// All products in the system
	// in: body
	Body data.Product
}

// No content is returned when product is updated or deleted
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}
