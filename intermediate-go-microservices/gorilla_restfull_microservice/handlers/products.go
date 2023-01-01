package handlers

import (
	"log"
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

type Products struct {
	tracer *log.Logger
}

func NewProducts(tracer *log.Logger) *Products {
	return &Products{tracer}
}

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

/*
restful post products method on Products handler struct
try $ curl -v localhost:5000/products -d '{}'
$ curl -v localhost:5000/products -d '{"id": 3 ,"name":"Something","description": "everything","Price": 0.00, "glaze": "nothing"}'
*/
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.tracer.Println("Handle POST Product")

	newProd := &data.Product{}
	err := newProd.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "failed to unmarshal json", http.StatusInternalServerError)
		p.tracer.Println(err)
	}

	/* %#v will show values & fields both*/
	p.tracer.Printf("NewProd: %#v", newProd)
	/*try $ curl -v localhost:5000/products -d '{"id": 10 ,"name":"Something","description": "everything","Price": 0.00, "glaze": "nothing"}'*/
	data.AddProduct(newProd)
}
