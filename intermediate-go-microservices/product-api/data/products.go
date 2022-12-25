package data

import "time"

/*
Donuts
Product defines the DTO/structure of the API product
*/
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Glaze       string  `json:"glaze"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

/*sample simulated Stubs, how a product list will look like in DB*/
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Choc-o-Choc",
		Description: "Made with bar one, Shell donut filled with dark chocolate",
		Price:       109.00,
		Glaze:       "chocolate",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Crispy Crunch",
		Description: "Shell donut frosted with dark compound",
		Price:       159.00,
		Glaze:       "choco-coated-crisples",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

/*Data access method to access product[donut list]*/
func GetProducts() []*Product {
	return productList
}
