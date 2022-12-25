package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	// create product
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}

	// put product
	if r.Method == http.MethodPut {
		// expects an id in the uri
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.tracer.Println("INVALID URI")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.tracer.Println("INVALID URI id access")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]

		id, err := strconv.Atoi(idString)
		if err != nil {
			p.tracer.Println("Unable to convert string to id", err)
			http.Error(rw, "Failed to parse id from uri", http.StatusInternalServerError)
			return
		}

		p.tracer.Println("grabbed ID", id)

	}

	//catch all other request method
	// try $ curl localhost:5000/products -XDELETE -v
	rw.WriteHeader(http.StatusNotImplemented)
}

/*
restful get products  method on Products handler struct
*/
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {

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
