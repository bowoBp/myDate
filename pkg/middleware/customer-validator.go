package middleware

import (
	"github.com/go-playground/validator/v10"
	"log"
	"strconv"
	"strings"
)

func validateNonZeroAndGreaterThan(fl validator.FieldLevel) bool {
	meter := fl.Field().Float()
	param := fl.Param()
	targetField := fl.Parent().FieldByName(param)
	t := targetField.Interface()
	tu, ok := t.(int)
	if !ok {
		return false
	}

	if meter > 0 {
		if tu > 0 {
			return true
		}
		return false
	}

	return true
}

func validateEnumValue(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if input == "" {
		return true
	}
	params := fl.Param()
	if !strings.Contains(params, input) {
		return false
	}

	return true
}

func validateMonthYearFormat(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	r := strings.Split(input, "-")
	if len(r) != 2 {
		return false
	}
	month, err := strconv.ParseUint(r[0], 10, 64)
	if err != nil {
		log.Println(err)
		return false
	}

	year, err := strconv.ParseUint(r[1], 10, 64)
	if err != nil {
		log.Println(err)
		return false
	}

	if (month > 12 && month < 1) &&
		(year > 2100 && year < 2000) {
		return false
	}

	return true
}
