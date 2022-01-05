package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/devtimx/GoSimpleApiRest/db"
	"github.com/devtimx/GoSimpleApiRest/jwt"
	"github.com/devtimx/GoSimpleApiRest/models"
)

/*Login, make the login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid User or Password"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email required", 400)
		return
	}
	doc, find := db.TryLogin(t.Email, t.Password)
	if find == false {
		http.Error(w, "Invalid User or Password ", 400)
		return
	}
	jwtkey, err := jwt.MakeJWT(doc)
	if err != nil {
		http.Error(w, "Error to try generate token "+err.Error(), 400)
		return
	}
	resp := models.ResponseLogin{
		Token: jwtkey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
