package model

import "github.com/golang-jwt/jwt/v5"

type RegisterInput struct {
	Username string `json:"username"  orm:"username"`
	Password string `json:"password"  orm:"password"`
	Email    string `json:"email"     orm:"email"   `
}

type JwtClaimsInput struct {
	UserId   string
	UserName string
	jwt.RegisteredClaims
}
