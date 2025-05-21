package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Преобразование ошибок валидации в map
func ValidationErrorsToMap(err error) map[string]string {
	errs := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errs[field] = "This field is required"
			case "email":
				errs[field] = "Invalid email format"
			case "min":
				errs[field] = "Value is too short"
			case "max":
				errs[field] = "Value is too long"
			default:
				errs[field] = "Invalid value"
			}
		}
	}

	return errs
}

// Проверка на нулевое значение структуры
func IsZero(value interface{}) bool {
	v := reflect.ValueOf(value)
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
