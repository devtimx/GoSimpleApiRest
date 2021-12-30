package db

import (
	"GoSimpleApiRest/models"
	"context"
	"fmt"
	"time"
)

/*CheckExistUser get an email and check if exist in the database */
func CheckExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `SELECT id,firstName,lastName,birthDate,email,password 
	FROM users WHERE email=$1;`
	db := PsqlCN.QueryRowContext(ctx, query, email)
	var u models.User
	err := db.Scan(&u.ID, &u.FirstName, &u.LastName, &u.BirthDate, &u.Email, &u.Password)
	ID := fmt.Sprint(u.ID)
	if err != nil {
		return u, false, ID
	}
	return u, true, ID

}
