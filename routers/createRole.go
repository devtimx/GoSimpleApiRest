package routers

import (
	"encoding/json"
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/db"
	"github.com/devtimx/GoSimpleApiRest/models"
)

func CreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var rl models.Role
	err := json.NewDecoder(r.Body).Decode(&rl)
	if err != nil {
		http.Error(w, "Failed to get data "+err.Error(), 400)
		return
	}

	if len(rl.Name) == 0 {
		http.Error(w, "Name required "+err.Error(), 400)
		return
	}

	status, err := db.InsertRole(rl)
	if err != nil {
		http.Error(w, "Error trying to register role "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Failed to insert role ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
