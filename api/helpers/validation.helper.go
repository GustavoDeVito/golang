package helpers

import "github.com/go-playground/validator/v10"

func formatterValidation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, "Field: "+e.StructField()+". Error: "+e.Tag())
	}

	return errors
}

func Validation(s interface{}) []string {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return formatterValidation(err)
	}

	return nil
}
