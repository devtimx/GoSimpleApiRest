package db

import (
	"context"
	"fmt"
	"time"

	"github.com/devtimx/GoSimpleApiRest/models"
)

var MyId string

/*CheckExistUser get an email and check if exist in the database */
func CheckExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `SELECT * FROM users WHERE email=$1;`
	db := PsqlCN.QueryRowContext(ctx, query, email)
	var u models.User
	err := db.Scan(&u.ID, &u.FirstName, &u.LastName, &u.BirthDate, &u.Email, &u.Password, &u.Profile, &u.Status, &u.CreatedAt, &u.UpdateAt)
	ID := fmt.Sprint(u.ID)
	if err != nil {
		return u, false, ID
	}
	MyId = ID
	return u, true, ID

}
