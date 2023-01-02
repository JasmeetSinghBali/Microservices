package data

import (
	"testing"
)

/*eventually fails as Product does not have name specified which is required*/
func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "grape-lime",
		Price: 11.99,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
