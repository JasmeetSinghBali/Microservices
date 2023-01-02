package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
	"github.com/gorilla/mux"
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

// Key: value based context http struct
type KeyProduct struct{}

/*
restful post products method on Products handler struct
try $ curl -v localhost:5000/products -d '{}'
$ curl -v localhost:5000/products -d '{"id": 3 ,"name":"Something","description": "everything","Price": 0.00, "glaze": "nothing"}'
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
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
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
