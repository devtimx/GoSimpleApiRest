package routers

import (
	"encoding/json"
	"net/http"

	"github.com/devtimx/GoSimpleApiRest/db"
	"github.com/devtimx/GoSimpleApiRest/models"
)

func CreatePermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var p models.Permission
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Failed to get data "+err.Error(), 400)
		return
	}

	if len(p.Name) == 0 {
		http.Error(w, "Name required "+err.Error(), 400)
		return
	}

	status, err := db.InsertPermission(p)
	if err != nil {
		http.Error(w, "Error trying to register permission "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Failed to insert permission ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
