package db

import (
	"GoSimpleApiRest/models"
	"context"
	"time"
)

/*InsertRecord */
func InsertRecord(u models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `INSERT INTO users(firstName,lastName,birthDate,email,password) values($1,$2,$3,$4,$5) RETURNING id;`

	u.Password, _ = HashPassword(u.Password)

	stmt, err := PsqlCN.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	db := stmt.QueryRowContext(ctx, u.FirstName, u.LastName, u.BirthDate, u.Email, u.Password)
	err = db.Scan(&u.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
