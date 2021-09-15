package routes

import (
	"golang-fiber-boilerplate/controllers"
	"golang-fiber-boilerplate/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/markbates/pkger"
);

func Setup (app *fiber.App) {
	api := app.Group("/api")

	/*
	 * Version 1
	 */
	v1 := api.Group("/v1", middlewares.RequestAuthencation)

	user := v1.Group("/user")
	user.Post("/create",controllers.CreateUser)
	user.Post("/login", controllers.AuthenticateUser)
	user.Get("/:id", controllers.UserList)
	
	cdn := app.Group("/cdn")
	cdn.Use("/images", filesystem.New(
		filesystem.Config{
			Root: pkger.Dir("/public/images"),
		},
	))
}