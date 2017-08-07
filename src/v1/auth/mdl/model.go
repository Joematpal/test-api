package authMdl

import (
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Session struct
type Session struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	jwt.StandardClaims
}
