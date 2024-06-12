package models

type User struct {
	ID       int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Record struct {
	ID      int `json:"id"`
	UserID  string `json:"user_id"`
	Name string `json:"name"`
	Content string `json:"content"`
	Cost int `json:"cost"`
}

type ResponseData struct {
    Users     []User   `json:"users"`
    Records   []Record `json:"records"`
    UserIDs   []int    `json:"user_ids"`
    RecordIDs []int    `json:"record_ids"`
    Message   string   `json:"message"`
}

type RequestData struct{
	
}