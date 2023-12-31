package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/j23063519/tech-school/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
