package models

import jwt "github.com/dgrijalva/jwt-go"

type Claim struct {
	Email string `json:"email"`
	ID    uint   `json:"id,omitempty"`
	jwt.StandardClaims
}
