package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

/*
Donuts
Product defines the DTO/structure of the API product
*/
type Product struct {
	// the id of the user
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	Glaze       string  `json:"glaze" validate:"required,glaze"`
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

// Products slice of type Product
type Products []*Product

/*
a method to the type Products named ToJSON
that encodes the data---> json via NewEncoder better performance than json.Marshall as it does not requires allocation for output buffer & overhead processing which is present with json.Marshal
*/
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}

func getRandomId() int {
	lengthProductList := productList[len(productList)-1]
	return lengthProductList.ID + 1
}

func AddProduct(p *Product) {
	p.ID = getRandomId()
	productList = append(productList, p)
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)

	if i == -1 {
		return ErrProductNotFound
	}

	productList[i] = &p

	return nil
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:i], productList[i+1])
	return nil
}
