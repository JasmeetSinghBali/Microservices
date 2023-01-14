package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// wrapping the validators fieldError
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// collection/slice of validation errors
type ValidationErrors []ValidationError

// convert slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation struct
type Validation struct {
	validate *validator.Validate
}

// create new validation type method on Validation struct
func NewValidation() *Validation {
	// ✨register the custom validation function to validate instance via passing
	// ✨name of the custom validation function & the function to perform the validation
	validate := validator.New()
	validate.RegisterValidation("glaze", validateGlaze)

	// store the mutated validate i.e glaze validation at Validation address struct
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

func validateGlaze(fl validator.FieldValue) bool {
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
