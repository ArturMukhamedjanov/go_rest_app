package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_rest_app/models"
	"net/http"
	"strconv"
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

type RestServerService struct {
	URL string
}

func NewRestServerService() *RestServerService {
	return &RestServerService{
		URL: "http://localhost:8000",
	}
}

var activeUserId int

func SetActiveUserId(id int) {
	activeUserId = id
}

func sendRequest(method HTTPMethod, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(string(method), url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	return resp, nil
}

func (s *RestServerService) RegisterUser(user models.User) (string, int) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return "error when serializes object", -1
	}
	resp, err := sendRequest(POST, s.URL+"/register", userJSON)
	if err != nil {
		return "failed to send request user", -1
	}
	if resp.StatusCode == http.StatusBadRequest {
		return "Wrong data", -1
	}
	if resp.StatusCode == http.StatusConflict {
		return "User with that username already exists", -1
	}
	if resp.StatusCode == http.StatusCreated {
		var result struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return "error when parsing response", -1
		}
		return "User successfully created", result.ID
	}
	return "unexpected server response", -1
}

func (s *RestServerService) AuthentificateUser(user models.User) int {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return -1
	}
	resp, err := sendRequest(GET, s.URL+"/authentificate", userJSON)
	if err != nil {
		return -1
	}
	if resp.StatusCode == http.StatusAccepted {
		var result struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return -1
		}
		return result.ID
	}
	return -1
}

func (s *RestServerService) AddRecord(record models.Record) string {
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return "error when serializes object"
	}
	resp, err := sendRequest(POST, s.URL+"/record", recordJSON)
	if err != nil {
		return "failed to send request record"
	}
	if resp.StatusCode == http.StatusBadRequest {
		return "Wrong data"
	}
	if resp.StatusCode == http.StatusInternalServerError {
		return "Internal server error"
	}
	if resp.StatusCode == http.StatusCreated {
		return "Record successfully created"
	}
	return "unexpected server response"
}

func (s *RestServerService) GetAllRecords() ([]models.Record, error) {
	resp, err := sendRequest(GET, s.URL+"/records/user/"+strconv.Itoa(activeUserId), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request for records by user ID: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected server response: %v", resp.Status)
	}
	var records []models.Record
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		return nil, fmt.Errorf("error when parsing response: %v", err)
	}
	return records, nil
}

func (s *RestServerService) GetRecordByID(recordID int) (*models.Record, error) {
	resp, err := sendRequest(GET, s.URL+"/records/user/"+strconv.Itoa(activeUserId)+"/record/"+strconv.Itoa(recordID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request for record by ID: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected server response: %v", resp.Status)
	}
	var record models.Record
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		return nil, fmt.Errorf("error when parsing response: %v", err)
	}
	return &record, nil
}

func (s *RestServerService) GetRecordsByName(name string) ([]models.Record, error) {
	resp, err := sendRequest(GET, s.URL+"/records/user/"+strconv.Itoa(activeUserId)+"/name/"+name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request for records by name: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected server response: %v", resp.Status)
	}
	var records []models.Record
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		return nil, fmt.Errorf("error when parsing response: %v", err)
	}
	return records, nil
}
