package service

import (
	"go_rest_app/models"
)

type UiService interface {
	GetUser() models.User
	GetRecord() models.Record
	GetUserID() int
	GetRecordID() int
	GetCommandName() string
	OutputUser(models.User) 
	OutputRecord(models.Record)
	OutputString(string)

}
