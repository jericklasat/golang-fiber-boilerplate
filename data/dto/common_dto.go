package dto

import "github.com/golang-jwt/jwt"

type Claims struct {
	Issuer string `json:"issuer"`
	jwt.StandardClaims
}