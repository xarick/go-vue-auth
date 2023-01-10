package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xarick/gin-sso/dto"
)

func GenerateJWT(JWTSecret string, JWTExpTime int, sub uint) (tokenString string, err error) {

	var JWTSecretByte = []byte(JWTSecret)
	if len(JWTSecret) < 1 {
		err = errors.New("secret key is empty")
		return
	}

	issuetionTime := time.Now()
	expirationTime := time.Now().Add(time.Duration(JWTExpTime) * time.Minute)

	claims := &dto.JWTClaim{
		Sub: sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			IssuedAt:  &jwt.NumericDate{Time: issuetionTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("okkkk")

	tokenString, err = token.SignedString(JWTSecretByte)
	return
}

func ValidateToken(signedToken string) (err error) {

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) < 1 {
		err = errors.New("secret key is empty")
		return
	}

	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*dto.JWTClaim)
	if !ok {
		err = errors.New("could not parse claims")
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
