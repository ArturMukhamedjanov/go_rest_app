package handlers

import (
	"encoding/json"
	"go_rest_app/models"
	"go_rest_app/server/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var repo repositories.Repository

func InitHandlers(repository repositories.Repository) {
	repo = repository
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !models.CheckCorrectUser(user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	existingUser, err := repo.GetUserByUsername(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	id, err := repo.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func AuthentificateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !models.CheckCorrectUser(user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authenticated := repo.AuthentificateUser(user)
	if authenticated{
		w.WriteHeader(http.StatusAccepted)
		exsistedUser, _ := repo.GetUserByUsername(user.Username)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": exsistedUser.ID,
		})
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func AddRecord(w http.ResponseWriter, r *http.Request){
	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !models.CheckCorrectRecord(record) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := repo.AddRecord(record)
	if err == nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func GetRecordsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	records := repo.GetRecordsByOwnerId(userID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records)
}

func GetRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordID, err := strconv.Atoi(vars["record_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	record := repo.GetRecordsById(recordID)
	if record.UserID != userID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(record)
}

func GetRecordsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordName := vars["name"]
	records := repo.GetRecordsByName(recordName)
	filteredRecords := []models.Record{}
	for _, record := range records {
		if record.UserID == userID {
			filteredRecords = append(filteredRecords, record)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filteredRecords)
}

