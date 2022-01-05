package middleware

import (
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/routers"
)

/*ValidJWT validate the token received in the request*/
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
