package middleware

import (
	"GoSimpleApiRest/routers"
	"net/http"
)

func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Autorization"))
		if err != nil {
			http.Error(w, "Error in the Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
