package handlers

import (
	"net/http"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product via id
//
// responses:
//		201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// restful delete product via id method on Products handler struct
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.tracer.Println("deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.tracer.Println("deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.tracer.Println("deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
