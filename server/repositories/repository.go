package repositories

import "go_rest_app/models"

type Repository interface {
	InitDB()
	CreateUser(user models.User) (int, error)
	GetUserByUsername(username string) (*models.User, error)
	DeleteUserByUsername(username string) error
	AuthentificateUser(user models.User) bool
	AddRecord(record models.Record) error
	GetRecordsByOwnerId(id int) []models.Record
	GetRecordsById(id int) models.Record
	GetRecordsByName(name string) []models.Record
}
