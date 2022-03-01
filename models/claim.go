package models

import jwt "github.com/dgrijalva/jwt-go"

/*Claim structure for processing the token */
type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id,omitempty"`
	jwt.StandardClaims
}
