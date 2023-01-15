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
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// single product returned in the response
// swagger:response productResponse
type productResponseWrapper struct {
	// All products in the system
	// in: body
	Body data.Product
}

// No content is returned when product is updated or deleted
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to be deleted from db
	// in:path
	// required:true
	ID int `json:"id"`
}
