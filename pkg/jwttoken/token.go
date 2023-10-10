package jwttoken

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type JWTToken struct {
	secretKey string
}

func New(secretKey string) *JWTToken {
	return &JWTToken{
		secretKey: secretKey,
	}
}

func (j *JWTToken) CreateToken(id string, email string, duration time.Duration) (tokenString string, err error) {
	expirationTime := time.Now().Add(duration)
	claims := &JWTClaim{
		UserID: id,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println("Secret key: ", j.secretKey)
	tokenString, err = token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Printf("Token sign error %s", err)
		return "", err
	}
	return tokenString, nil
}

func (j *JWTToken) ValidateToken(signedToken string) (claim *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}
