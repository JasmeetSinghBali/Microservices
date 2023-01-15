package handlers

import (
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//		200: productsResponse

// restful get all products method on Products handler struct
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.tracer.Println("Handle GET Products")

	/*access the method in data package to get product list*/
	listOfProducts := data.GetProducts()

	/*
		calls ToJSON , converts data ---> slice of byte to return it to client [marshalling]
	*/
	err := data.ToJSON(listOfProducts, rw)
	if err != nil {
		http.Error(rw, "Failed to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingle
// Returns a single product from DB
// responses:
//		200: productResponse
//		404: errorResponse

// restful get product via id  method on Products handler struct
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
