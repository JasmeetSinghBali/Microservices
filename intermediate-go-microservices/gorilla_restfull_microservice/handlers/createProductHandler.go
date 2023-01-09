package handlers

import (
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
// 422: errorValidation
// 501: errorResponse

//restful post products method on Products handler struct
/*fails - try $ curl localhost:5000/ -X POST -d '{"Name": "New Donut"}'
pass - try  $ curl localhost:5000/ -X POST -d '{"Name": "New Donut", "Price": 7.99, "Glaze": "strawberry-bottom-vanilla"}'
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
