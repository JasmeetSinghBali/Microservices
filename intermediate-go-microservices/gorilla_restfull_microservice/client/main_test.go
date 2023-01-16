package main

import (
	"fmt"
	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/client/client"
	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/gorilla_restfull_microservice/client/client/products"
	"testing"
)

// 📝 make sure the server is running, by default it dials/makes request to [::1]:80:
// while the go server is running on 5000, to change/over-ride the default dial target
// ✨ reff: gorilla_restapi_client.go in client/client/ dir WithHost & TransportConfig
func TestGeneratedClient(t *testing.T) {

	// overiding transport config to change the dialed target via http client to go server
	cfg := client.DefaultTransportConfig().WithHost("localhost:5000")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	// ref: products_client.go inside client/client generated http client
	// ✨calls ListProducts method on Products with params from NewListProductParams
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	// printing out response products from /products api call tested ✔
	fmt.Printf("%#v", prod.GetPayload()[0])
	fmt.Printf("%#v", prod.GetPayload()[1])
	t.Fail()
}
