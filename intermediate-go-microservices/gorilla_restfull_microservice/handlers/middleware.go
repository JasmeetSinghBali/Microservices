package handlers

import (
	"context"
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

/*
Middleware that validates the payload while creating or updating product
*/
func (p *Products) ProductValidationMiddleware(next http.Handler) http.Handler {
	// response is just an interface, where request is reff to the actual request passed on from client
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prd := data.Product{}

		err := data.FromJSON(prd, r.Body)
		if err != nil {
			p.tracer.Println("Failed to parse json & go oriented data/struct", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// 📝 validate/sanitize the product paylaod via Validate method exposed via data package reff: data/products.go
		errs := p.checker.Validate(prd)
		if len(errs) > 0 {
			p.tracer.Println("validating product errors", errs)

			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
		}

		// attaches the product i.e prd the unmarshalled json extracted from request body as newContext on the request object that is common and has reff from client--->middlewares--->handlers--->data access methods
		ctx := context.WithValue(r.Context(), KeyProduct{}, prd)
		// create new req with old common r (req) object reff with attached newContext
		req := r.WithContext(ctx)

		// pass control the http handler with mutated req & interface only rw
		next.ServeHTTP(rw, req)

	})
}
