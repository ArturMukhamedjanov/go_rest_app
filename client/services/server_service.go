package service

import "go_rest_app/models"

type ServerService interface {
	RegisterUser(models.User) (string, int)
	AuthentificateUser(models.User) (int)
	AddRecord(models.Record) string
	GetAllRecords() ([]models.Record, error)
	GetRecordByID(recordID int) (*models.Record, error)
	GetRecordsByName(name string) ([]models.Record, error)
}
