package middleware

import (
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/db"
)

func CheckPermission(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckPermission() == 0 {
			http.Error(w, "permission denied", 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
