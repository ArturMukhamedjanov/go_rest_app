package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func CheckCorrectUser(user User) bool {
	if user.Username == "" || user.Password == ""{
		return false
	}
	return true
}

type Record struct {
	ID      int    `json:"id"`
	UserID  int `json:"user_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Cost    int    `json:"cost"`
}

func CheckCorrectRecord(record Record) bool {
	if record.Content == "" || record.Name == "" || record.UserID == 0{
		return false
	}
	return true
}
