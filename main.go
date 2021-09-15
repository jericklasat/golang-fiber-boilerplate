package main

import (
	database "golang-fiber-boilerplate/configs"
	"golang-fiber-boilerplate/repositories"
	"golang-fiber-boilerplate/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
);

func main() {
	app := fiber.New();
	db := database.Init();
	database.Migrate();

	repositories.DB = db;

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*",
	}));
	
	routes.Setup(app);
	
	app.Listen(":5000");
}