package dto

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}
