package main

import (
	"log"

	"github.com/devtimx/GoSimpleApiRest/db"
	"github.com/devtimx/GoSimpleApiRest/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Failed connection with the database")
		return
	}
	handlers.Handlers()
}
