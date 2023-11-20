package jwttoken

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
