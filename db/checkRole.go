package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/devtimx/GoSimpleApiRest/models"
)

func CheckRole() (int, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Println("MyId: " + MyId)
	query := `SELECT profile FROM users WHERE id=$1;`
	db := PsqlCN.QueryRowContext(ctx, query, MyId)
	var u models.User
	err := db.Scan(&u.Profile)
	Profile := fmt.Sprint(u.Profile)
	if err != nil {
		return 0, ""
	}
	return 1, Profile
}
