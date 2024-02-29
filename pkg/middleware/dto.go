package middleware

import "github.com/golang-jwt/jwt/v4"

type DefaultUserClaim struct {
	UserData UserData `json:"userData"`
	jwt.RegisteredClaims
}

type UserData struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
}
