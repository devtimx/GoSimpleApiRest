package db

import (
	"context"
	"strings"
	"time"

	"github.com/devtimx/GoSimpleApiRest/models"
)

func InsertRole(r models.Role) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `INSERT INTO roles(name,guard_name,created_at,updated_at) values($1,$2,$3,$4) RETURNING id;`

	stmt, err := PsqlCN.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	r.GuardName = "web"
	db := stmt.QueryRowContext(ctx, strings.ToUpper(r.Name), r.GuardName, time.Now(), time.Now())
	err = db.Scan(&r.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
