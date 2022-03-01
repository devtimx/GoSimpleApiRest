package routers

import (
	"encoding/json"
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/db"
	"github.com/devtimx/GoSimpleApiRest/models"
)

/*SignUp create user records in the database*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Failed to get data "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email required "+err.Error(), 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password must be at least 6 characters "+err.Error(), 400)
		return
	}
	_, found, _ := db.CheckExistUser(t.Email)
	if found == true {
		http.Error(w, "The email exists in the database "+err.Error(), 400)
		return
	}

	status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "Error trying to register user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Failed to insert user ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
