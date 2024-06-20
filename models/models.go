package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type Record struct {
	ID      int    `json:"id"`
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Cost    int    `json:"cost"`
}