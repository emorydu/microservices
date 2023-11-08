package data

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidationError wraps the validators FieldError, so we do not
// expose this to out code
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Filed Validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//
//	if ve, ok := err.(validator.ValidationErrors); ok {
//				fmt.Println(ve.Namespace())
//				fmt.Println(ve.Field())
//				fmt.Println(ve.StructNamespace())
//				fmt.Println(ve.StructField())
//				fmt.Println(ve.Tag())
//				fmt.Println(ve.ActualTag())
//				fmt.Println(ve.Kind())
//				fmt.Println(ve.Type())
//				fmt.Println(ve.Value())
//				fmt.Println(ve.Param())
//				fmt.Println()
//		}
func (v *Validation) Validate(i any) ValidationErrors {
	var errs validator.ValidationErrors
	errors.As(v.validate.Struct(i), &errs)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		// cats the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// NewValidation create a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate: validate}
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format xxx-xxx-xxx
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
