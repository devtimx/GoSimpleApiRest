package jwt

import (
	"time"

	"github.com/devtimx/GoSimpleApiRest/models"

	jwt "github.com/dgrijalva/jwt-go"
)

/*MakeJWT make the encode with jwt*/
func MakeJWT(t models.User) (string, error) {
	myKey := []byte("GoSimpleApiRest")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"firstName": t.FirstName,
		"lastName":  t.LastName,
		"birthDate": t.BirthDate,
		"profile":   t.Profile,
		"status":    t.Status,
		"id":        t.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
