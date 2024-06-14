package repositories

import "go_rest_app/models"

type Repository interface {
	InitDB()
	CreateUser() (string, error)
	GetUserByUsername() models.User
	DeleteUserByUsername() string
	
}
