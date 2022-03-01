package middleware

import (
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/db"
)

func CheckRole(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err, profile := db.CheckRole()
		if err == 0 && profile == "" {
			http.Error(w, "permission denied", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
