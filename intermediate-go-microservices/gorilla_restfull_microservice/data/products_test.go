package data

import (
	"testing"
)

/*eventually fails as Product does not have name specified which is required*/
func TestCheckValidation(t *testing.T) {
	p := &Product{}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
