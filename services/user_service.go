package services

import (
	"encoding/json"
	"golang-fiber-boilerplate/data/dto"
	"golang-fiber-boilerplate/data/response"
	"golang-fiber-boilerplate/enums"
	"golang-fiber-boilerplate/repositories"
	"golang-fiber-boilerplate/validations"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *dto.User) response.DataResponse {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	baseData := dto.BaseDto{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userData := dto.User{
		Name:       user.Name,
		Email:      user.Email,
		Password:   string(password),
		IsVerified: user.IsVerified,
		Level:      user.Level,
		Status:     user.Status,
		BaseDto:    baseData,
	}

	validations.UserStruct = userData

	errors := validations.ValidateUserStruct()
	if errors != nil {
		return response.DataResponse{
			Status:  fiber.StatusNotAcceptable,
			Message: "Input error",
			Error:   errors,
		}
	}

	repositoryResponse := repositories.CreateUser(userData)
	if repositoryResponse != "OK" {
		return response.DataResponse{
			Status:  enums.RepositoryError,
			Message: repositoryResponse,
		}
	}

	data, _ := json.Marshal(userData)

	return response.DataResponse{
		Status:  fiber.StatusOK,
		Message: "Creation successful",
		Data:    string(data),
	}
}

func AuthenticateUser(credentials map[string]string) response.DataResponse {
	userData := repositories.GetUserByEmail(credentials["email"])
	if userData.ID == 0 {
		return response.DataResponse{
			Status:  fiber.StatusNotFound,
			Message: "Account not found.",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(credentials["password"])); err != nil {
		return response.DataResponse{
			Status:  fiber.StatusNotFound,
			Message: "Incorrect password.",
		}
	}

	claims := dto.Claims{
		Issuer: strconv.Itoa(int(userData.ID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return response.DataResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Cannot generate token.",
		}
	}

	return response.DataResponse{
		Status:  fiber.StatusOK,
		Message: "Authentication successful.",
		Data:    signedString,
	}
}
