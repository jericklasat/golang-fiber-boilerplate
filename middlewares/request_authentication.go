package middlewares

import (
	"golang-fiber-boilerplate/helpers"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func RequestAuthencation(ctx *fiber.Ctx) error {
	ctx.Set("X-XSS-Protection", "1; mode=block");
	ctx.Set("Strict-Transport-Security", "max-age=5184000");
	ctx.Set("X-DNS-Prefetch-Control", "off");

	urlWhitelist := []string {
		"/api/v1/user/login",
		"/api/v1/user/create",
	}

	// Check if url is in whitelist
	if helpers.InArray(ctx.OriginalURL(), urlWhitelist) || strings.Contains(ctx.OriginalURL(), "cdn") {
		return ctx.Next();
	}

	authorizationToken  := string(ctx.Request().Header.Peek("Authorization"));

	// Check if user passed authorization header.
	if authorizationToken == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}

	// Check if token is valid
	_, err := jwt.ParseWithClaims(authorizationToken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	});

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}

	return ctx.Next()
}