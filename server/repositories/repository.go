package repositories

import "go_rest_app/models"

type Repository interface {
	InitDB()
	CreateUser(user models.User) (string, error)
	GetUserByUsername(username string)  (*models.User, error)
	DeleteUserByUsername(username string) string
}
