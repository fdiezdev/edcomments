package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim -> Auth tokens
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
