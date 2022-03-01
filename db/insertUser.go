package db

import (
	"context"
	"strings"
	"time"

	"github.com/devtimx/GoSimpleApiRest/models"
)

/*InsertRecord insert data to database*/
func InsertUser(u models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `INSERT INTO users(first_name,last_name,birth_date,email,password,profile,status,created_at,updated_at) values($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id;`

	u.Password, _ = HashPassword(u.Password)

	stmt, err := PsqlCN.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	db := stmt.QueryRowContext(ctx, u.FirstName, u.LastName, u.BirthDate, strings.TrimSpace(u.Email), u.Password, strings.ToUpper(u.Profile), true, time.Now(), time.Now())
	err = db.Scan(&u.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
