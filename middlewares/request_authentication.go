package middlewares

import (
	"golang-fiber-boilerplate/data/dto"
	"golang-fiber-boilerplate/helpers"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func RequestAuthentication(ctx *fiber.Ctx) error {
	ctx.Set("X-XSS-Protection", "1; mode=block");
	ctx.Set("Strict-Transport-Security", "max-age=5184000")
	ctx.Set("X-DNS-Prefetch-Control", "off")

	urlWhitelist := getWhiteList()

	if helpers.InArray(ctx.OriginalURL(), urlWhitelist) || strings.Contains(ctx.OriginalURL(), "cdn") {
		return ctx.Next()
	}

	authorizationToken := getAuthorizationToken(ctx)

	if authorizationToken == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	_, err := jwt.ParseWithClaims(authorizationToken, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}

func getAuthorizationToken(ctx *fiber.Ctx) string {
	authorizationToken := string(ctx.Request().Header.Peek("Authorization"))
	authorizationToken = strings.Replace(authorizationToken, "Bearer ", "", 1)

	return authorizationToken
}

func getWhiteList() []string {

	return []string {
		"/api/v1/user/login",
		"/api/v1/user/create",
	}
}