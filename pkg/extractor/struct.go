package extractor

import "github.com/golang-jwt/jwt/v4"

type Token struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}
