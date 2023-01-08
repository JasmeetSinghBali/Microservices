package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

/*
Donuts
Product defines the DTO/structure of the API product
*/
type Product struct {
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

/*
a method to the type Product named FromJSON
to decode json ----> data via NewDecoder
*/
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

/*validation method*/
func (p *Product) Validate() error {
	validate := validator.New()
	// ✨register the custom validation function to validate instance via passing
	// ✨name of the custom validation function & the function to perform the validation
	validate.RegisterValidation("glaze", validateGlaze)
	return validate.Struct(p)
}

func validateGlaze(fl validator.FieldLevel) bool {

	// glaze format flavour(any)-top/bottom-filling(any)
	re := regexp.MustCompile(`[a-z]+-(top|bottom)+-[a-z]+`)

	// grab the field value and convert to string and then compare with glaze format regex
	// matches return matched slice of string
	matches := re.FindAllString(fl.Field().String(), -1)

	// if no match found
	if len(matches) < 1 {
		return false
	}
	return true
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

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}
