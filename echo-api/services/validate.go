package services

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(value interface{}) []map[string]string {
	validate := validator.New()
	// custom field tag to json tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(value); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		errors := make([]map[string]string, len(validateErrors))
		for i, e := range validateErrors {
			errors[i] = map[string]string{"field": e.Field(), "msg": msgForTag(e.Tag())}
		}
		return errors
	}
	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return "This field is invalid"

}
