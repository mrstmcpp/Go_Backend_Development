package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

//centralized validator obj
func InitValidator() {
	Validate = validator.New()
}
