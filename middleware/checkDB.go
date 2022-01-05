package middleware

import (
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/db"
)

/*CheckDB check the status of the connection with the database*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "lost connection with the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
