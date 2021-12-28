package main

import (
	"GoSimpleApiRest/db"
	"GoSimpleApiRest/handlers"
	"log"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Failed connection with the database")
		return
	}
	handlers.Handlers()
}
