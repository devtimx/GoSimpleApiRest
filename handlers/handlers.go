package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/devtimx/GoSimpleApiRest/middleware"
	"github.com/devtimx/GoSimpleApiRest/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers, set port and put the server to listen, enable route functions  */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middleware.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/role", middleware.CheckDB(middleware.ValidJWT(routers.CreateRole))).Methods("POST")
	router.HandleFunc("/permission", middleware.CheckDB(middleware.ValidJWT(middleware.CheckRole(routers.CreatePermission)))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
