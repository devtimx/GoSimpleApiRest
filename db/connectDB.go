package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var PsqlCN = ConnectDB()

/*ConnectDB this function connect the database with the application*/
func ConnectDB() *sql.DB {
	connStr := "postgres://debian11:toor@localhost:5432/gotest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return db
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return db
	}
	log.Println("Success connection with database")
	return db
}

/*CheckConnection this function check if exist communication with the database */
func CheckConnection() int {
	err := PsqlCN.Ping()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
