package handlers

import (
	"encoding/json"
	"go_rest_app/models"
	"go_rest_app/server/repositories"
	"net/http"
)

var repo repositories.Repository

func InitHandlers(repository repositories.Repository) {
	repo = repository
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	repo.CreateUser()
}
