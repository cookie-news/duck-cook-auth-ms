package entity

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
	jwt.RegisteredClaims
}
