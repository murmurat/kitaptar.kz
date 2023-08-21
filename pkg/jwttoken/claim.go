package jwttoken

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	//Username string `json:"username"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
