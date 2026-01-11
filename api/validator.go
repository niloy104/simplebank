package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/niloy104/simplebank/util"
)

var validCurrency validator.Func = func(feildlevel validator.FieldLevel) bool {
	if currency, ok := feildlevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
