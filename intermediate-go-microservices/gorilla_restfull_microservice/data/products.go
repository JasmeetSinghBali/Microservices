package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

/*
Donuts
Product defines the DTO/structure of the API product
*/
// swagger:model
type Product struct {

	// the id of the user
	//
	// required: true
	// min: 1
	ID int `json:"id"` // unique indentifier for the product

	// the name for the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	//the description for the product
	//
	// required: false
	// maxLength: 10000
	Description string `json:"description"`

	//the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	//the glaze for the product/donut
	//
	// required: true
	// pattern: [a-z]+-(top|bottom)+-[a-z]+
	Glaze     string `json:"glaze" validate:"required,glaze"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
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
