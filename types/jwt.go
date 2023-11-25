package types

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	ID     uint   `json:"_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}
