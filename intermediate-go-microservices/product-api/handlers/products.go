package handlers

import (
	"log"
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/product-api/data"
)

type Products struct {
	tracer *log.Logger
}

func NewProducts(tracer *log.Logger) *Products {
	return &Products{tracer}
}

/*
controller/handler , that calls restful routes a/c to the method types
*/
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//catch all other request method
	// try $ curl localhost:5000/products -XDELETE -v
	rw.WriteHeader(http.StatusNotImplemented)
}

/*
restful get products  method on Products handler struct
*/
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
