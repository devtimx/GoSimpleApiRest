package db

import (
	"github.com/devtimx/GoSimpleApiRest/models"

	"golang.org/x/crypto/bcrypt"
)

/*TryLogin check password with password in the database*/
func TryLogin(email string, password string) (models.User, bool) {
	u, find, _ := CheckExistUser(email)
	if find == false {
		return u, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return u, false
	}
	return u, true
}
