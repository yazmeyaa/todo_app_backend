package services

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.MapClaims
	UserId int
}

type JWTService interface {
	Sign(payload UserClaims) (string, error)
	Verify(token string) bool
	Decode(token string) (*UserClaims, error)
}
