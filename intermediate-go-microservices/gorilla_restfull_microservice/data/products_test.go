package data

import (
	"testing"
)

/*eventually fails as Product does not have name specified which is required*/
func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "grape-lime",
		Price: 11.99,
		Glaze: "grapevine-top-orange",
		// 🧩 try dumb or any input that does not matches
		// format flavour(any i.e a-z)-top/bottom-filling(any i.e a-z)
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
