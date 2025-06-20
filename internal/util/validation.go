package util

import "github.com/go-playground/validator/v10"

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {

		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = v.Error()
		}
	}
	return res
}
