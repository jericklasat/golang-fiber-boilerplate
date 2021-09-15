package validations

import (
	"golang-fiber-boilerplate/data/dto"
	"golang-fiber-boilerplate/data/response"

	"github.com/go-playground/validator"
)
var UserStruct dto.User

func ValidateUserStruct() response.ErrorStruct {
	var returnErrors response.ErrorStruct

	if (dto.User{}) != UserStruct {
		validate := validator.New()
		err := validate.Struct(UserStruct)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				errorData := response.ErrorData {
					Field: err.StructField(),
					Tag: err.Tag(),
				};
				returnErrors = append(returnErrors, errorData)
			}
		}
	}
	return returnErrors
}