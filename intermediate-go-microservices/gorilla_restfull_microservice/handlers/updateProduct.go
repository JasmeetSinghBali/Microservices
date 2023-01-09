package handlers

import (
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

// swagger:route PUT /products products updateProduct
// Update a product via id
//
// responses:
// 	201: noContentReponse
// 404: errorResponse
// 422: errorValidation

// restful update product method on Products handler struct
// try curl localhost:5000/1 -X PUT -d '{"name": "Choc-o-Moc"}'
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	// 📝 Grab the unmarshalled product from KeyProduct struct context and store in prod
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.tracer.Println("Handle PUT Products", prod.ID)

	// make the method call to update product by extracetd id in data access layer products.go
	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.tracer.Println("product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found"}, rw)
		return
	}

	// writes the no content success header
	rw.WriteHeader(http.StatusNoContent)

}
