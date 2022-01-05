package routers

import (
	"errors"
	"strings"

	"github.com/devtimx/GoSimpleApiRest/db"
	"github.com/devtimx/GoSimpleApiRest/models"

	jwt "github.com/dgrijalva/jwt-go"
)

/*Email global data*/
var Email string

/*IDUser global data*/
var IDUser string

/*ProcessToken, processing for extract data of token*/
func ProcessToken(tk string) (models.Claim, bool, string, error) {
	myKey := []byte("GoSimpleApiRest")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return *claims, false, string(""), errors.New("Invalid token format")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find, _ := db.CheckExistUser(claims.Email)
		if find == true {
			Email = claims.Email
			IDUser = claims.Id
		}
		return *claims, find, IDUser, nil
	}
	if !tkn.Valid {
		return *claims, false, string(""), errors.New("invalid token")
	}
	return *claims, false, string(""), err
}
