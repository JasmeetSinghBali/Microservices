package handlers

import (
	"encoding/json"
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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	/*access the method in data package to get product list*/
	listOfProducts := data.GetProducts()

	/*convert data ---> slice of byte to return it to client [marshalling]*/
	d, err := json.Marshal(listOfProducts)
	if err != nil {
		http.Error(rw, "Failed to marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
