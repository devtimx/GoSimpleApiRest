package db

import (
	"context"
	"strings"
	"time"

	"github.com/devtimx/GoSimpleApiRest/models"
)

func InsertPermission(p models.Permission) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `INSERT INTO permissions(name,guard_name,created_at,updated_at) values($1,$2,$3,$4) RETURNING id;`

	stmt, err := PsqlCN.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	p.GuardName = "web"
	db := stmt.QueryRowContext(ctx, strings.ToUpper(p.Name), p.GuardName, time.Now(), time.Now())
	err = db.Scan(&p.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
