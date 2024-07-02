package services

import (
	"reflect"
	"strconv"
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

/*
Check this floating value have the right decimals

expect: decimal length that value should equal or less than

value: floating value
*/
func ValidDecimalsPlace(expect int, value float64) bool {
	actual := countDecimalPlaces(value)
	return actual <= expect
}

func countDecimalPlaces(num float64) int {
	// Convert float64 to string
	// ref: https://stackoverflow.com/a/76780465
	str := strconv.FormatFloat(num, 'E', -1, 64)

	// Split the string by decimal point
	parts := strings.Split(str, ".")

	// If there are two parts, count the length of the second part
	if len(parts) == 2 {
		return len(parts[1])
	}

	// Otherwise, there are no decimal places
	return 0
}
