package controllers

import (
	"golang-fiber-boilerplate/data/dto"
	"golang-fiber-boilerplate/data/response"
	"golang-fiber-boilerplate/services"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	user := new(dto.User)
	if err := ctx.BodyParser(user); err != nil {
		return err;
	}

	createUser := services.CreateUser(user)
	return ctx.JSON(createUser)
}

func AuthenticateUser(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err;
	}

	authenticateUser := services.AuthenticateUser(data)
	return ctx.JSON(authenticateUser)
}

func UserList(ctx *fiber.Ctx) error {
	dataResponse := response.DataResponse {
		Status: 200,
		Message: ctx.Params("id"),
	}
	return ctx.JSON(dataResponse)
}