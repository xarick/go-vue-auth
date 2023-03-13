package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xarick/gin-sso/dto"
	"github.com/xarick/gin-sso/initializers"
)

func GenerateJWT(sub uint) (tokenString string, err error) {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		err = errors.New("could not load environment variables")
		return
	}

	var JWTSecretByte = []byte(config.JWTSecret)
	if len(JWTSecretByte) < 1 {
		err = errors.New("secret key is empty")
		return
	}

	issuetionTime := time.Now()
	expirationTime := time.Now().Add(time.Duration(config.JWTExpTime) * time.Minute)

	claims := &dto.JWTClaim{
		Sub: sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			IssuedAt:  &jwt.NumericDate{Time: issuetionTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(JWTSecretByte)
	return
}

func ValidateToken(signedToken string) (err error) {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		err = errors.New("could not load environment variables")
		return
	}

	var jwtKey = []byte(config.JWTSecret)
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

func TokenToUserID(tokenStr string) (userID float64, err error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		log.Panic(err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID = claims["sub"].(float64)
	}
	return
}
