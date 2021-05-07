package services

import (
	"bnw-backend/data/dto"
	"bnw-backend/data/response"
	"bnw-backend/repositories"
	"bnw-backend/validations"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *dto.User) response.DataResponse {
	password,_ := bcrypt.GenerateFromPassword([]byte(user.Password), 14);
	base_data := dto.BaseDto {
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	};
	user_data := dto.User {
		Name: user.Name,
		Email: user.Email,
		Password: string(password),
		IsVerified: user.IsVerified,
		Level: user.Level,
		Status: user.Status,
		BaseDto: base_data,
	};
	
	// Set Validation Struct
	validations.UserStruct = user_data;
	// Validate
	errors := validations.ValidateUserStruct();
	if errors != nil {
		response := response.DataResponse {
			Status: fiber.StatusNotAcceptable,
			Message: "Input error",
			Error: errors,
		}
		return response;
	}

	// Handle repositories error
	repository_response := repositories.CreateUser(user_data);
	if (repository_response != "OK") {
		response := response.DataResponse {
			Status: 1062,
			Message: repository_response,
		}
		return response;
	}

	data,_ := json.Marshal(user_data);

	response := response.DataResponse {
		Status: fiber.StatusOK,
		Message: "Creation successful",
		Data: string(data),
	}
	return response;
}

func AuthenticateUser(credentials map[string]string) response.DataResponse {
	user_data := repositories.GetUserByEmail(credentials["email"]);
	if user_data.ID == 0 {
		response := response.DataResponse {
			Status: fiber.StatusNotFound,
			Message: "Account not found.",
		}
		return response;
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(user_data.Password), []byte(credentials["password"])); err != nil {
		response := response.DataResponse {
			Status: fiber.StatusNotFound,
			Message: "Incorrect password.",
		}
		return response;
	}
	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user_data.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	});

	token, err := claims.SignedString([]byte(os.Getenv("TOKEN_SECRET")));
	if err != nil {
		response := response.DataResponse {
			Status: fiber.StatusInternalServerError,
			Message: "Cannot generate token.",
		}
		fmt.Println(err);
		return response;
	}

	response := response.DataResponse {
		Status: fiber.StatusOK,
		Message: "Authentication successful.",
		Data: token,
	}
	return response;
}