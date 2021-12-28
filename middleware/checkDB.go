package middleware

import (
	"GoSimpleApiRest/db"
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "lost connection with the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
