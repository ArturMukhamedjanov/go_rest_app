package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_rest_app/models"
	"io"
	"net/http"
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

func NewRestServerService() RestServerService {
	return RestServerService{
		URL: "localhost:3000",
	}
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

func (s *RestServerService) RegisterUser(user models.User) string {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return "error when serializes object"
	}
	resp, err := sendRequest(POST, s.URL+"/register", userJSON)
	if err != nil {
		return "failed to register user"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("failed to read response body: %v", err)
	}

	return string(body)
}
