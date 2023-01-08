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
// swagger:meta
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

/*only for swagger doc response struct usecase*/
// A list of products returns in the response
// swagger:response productResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

type Products struct {
	tracer *log.Logger
}

type GenericError struct {
	Message string `json:"message"`
}

func NewProducts(tracer *log.Logger) *Products {
	return &Products{tracer}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

/*
restful get products  method on Products handler struct
*/
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.tracer.Println("Handle GET Products")

	/*access the method in data package to get product list*/
	listOfProducts := data.GetProducts()

	/*
		calls ToJSON , converts data ---> slice of byte to return it to client [marshalling]
	*/
	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed to marshal json", http.StatusInternalServerError)
	}
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
restful get product via id  method on Products handler struct
*/
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.tracer.Println("Handle GET Product by ID", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.tracer.Println("Error fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.tracer.Println("fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		p.tracer.Println("error in serializing product data to json", err)
	}

}

// Key: value based context http struct
type KeyProduct struct{}

/*
restful post products method on Products handler struct
fails - try $ curl localhost:5000/ -X POST -d '{"Name": "New Donut"}'
pass - try  $ curl localhost:5000/ -X POST -d '{"Name": "New Donut", "Price": 7.99, "Glaze": "strawberry-bottom-vanilla"}'^C
'
*/
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.tracer.Println("Handle POST Product")

	// 📝 Grab the unmarshalled product from KeyProduct struct context and store in newProd
	newProd := r.Context().Value(KeyProduct{}).(data.Product)

	/* %#v will show values & fields both*/
	p.tracer.Printf("NewProd: %#v", newProd)
	/*try $ curl -v localhost:5000/products -d '{"id": 10 ,"name":"Something","description": "everything","Price": 0.00, "glaze": "nothing"}'*/
	data.AddProduct(&newProd)
}

/*
restful update product method on Products handler struct
try curl localhost:5000/1 -X PUT -d '{"name": "Choc-o-Moc"}'
*/
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	// 📝 mux.Vars(r) is a map where r is http Request will hold the extracted uri/param provided while defining subrouter at upper level
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "failed to cast id string to int in request param url", http.StatusBadRequest)
		return
	}

	p.tracer.Println("Handle PUT Products", id)

	// 📝 Grab the unmarshalled product from KeyProduct struct context and store in prod
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// make the method call to update product by extracetd id in data access layer products.go
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.tracer.Println(err)
		http.Error(rw, "failed to update product", http.StatusInternalServerError)
		return
	}
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
