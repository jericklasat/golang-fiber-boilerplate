package services

import (
	"encoding/json"
	"fmt"
	"golang-fiber-boilerplate/data/dto"
	"golang-fiber-boilerplate/data/response"
	"golang-fiber-boilerplate/repositories"
	"golang-fiber-boilerplate/validations"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *dto.User) response.DataResponse {
	password,_ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	baseData := dto.BaseDto {
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userData := dto.User {
		Name: user.Name,
		Email: user.Email,
		Password: string(password),
		IsVerified: user.IsVerified,
		Level: user.Level,
		Status: user.Status,
		BaseDto: baseData,
	};
	
	// Set Validation Struct
	validations.UserStruct = userData
	// Validate
	errors := validations.ValidateUserStruct()
	if errors != nil {
		dataResponse := response.DataResponse {
			Status: fiber.StatusNotAcceptable,
			Message: "Input error",
			Error: errors,
		}
		return dataResponse
	}

	// Handle repositories error
	repositoryResponse := repositories.CreateUser(userData)
	if repositoryResponse != "OK" {
		dataResponse := response.DataResponse {
			Status: 1062,
			Message: repositoryResponse,
		}
		return dataResponse
	}

	data,_ := json.Marshal(userData)

	dataResponse := response.DataResponse {
		Status: fiber.StatusOK,
		Message: "Creation successful",
		Data: string(data),
	}
	return dataResponse
}

func AuthenticateUser(credentials map[string]string) response.DataResponse {
	userData := repositories.GetUserByEmail(credentials["email"])
	if userData.ID == 0 {
		dataResponse := response.DataResponse {
			Status: fiber.StatusNotFound,
			Message: "Account not found.",
		}
		return dataResponse
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(credentials["password"])); err != nil {
		dataResponse := response.DataResponse {
			Status: fiber.StatusNotFound,
			Message: "Incorrect password.",
		}
		return dataResponse
	}
	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(userData.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("TOKEN_SECRET")));
	if err != nil {
		dataResponse := response.DataResponse {
			Status: fiber.StatusInternalServerError,
			Message: "Cannot generate token.",
		}
		fmt.Println(err)
		return dataResponse
	}

	dataResponse := response.DataResponse {
		Status: fiber.StatusOK,
		Message: "Authentication successful.",
		Data: token,
	}
	return dataResponse
}