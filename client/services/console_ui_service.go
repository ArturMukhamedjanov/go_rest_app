package service

import (
	"fmt"
	"go_rest_app/models"
)

type ConsoleUiService struct{}

func (s *ConsoleUiService) GetUser() models.User {
	var user models.User
	fmt.Print("Enter username: ")
	fmt.Scan(&user.Username)
	fmt.Print("Enter password: ")
	fmt.Scan(&user.Password)
	return user
}

func (s *ConsoleUiService) GetRecord() models.Record {
	var record models.Record
	fmt.Print("Enter record name: ")
	fmt.Scan(&record.Name)
	fmt.Print("Enter record data: ")
	fmt.Scan(&record.Content)
	fmt.Print("Enter record cost: ")
	fmt.Scan(&record.Cost)
	return record
}

func (s *ConsoleUiService) GetUserID() int {
	var userID int
	fmt.Print("Enter user ID: ")
	fmt.Scan(&userID)
	return userID
}

func (s *ConsoleUiService) GetRecordID() int {
	var recordID int
	fmt.Print("Enter record ID: ")
	fmt.Scan(&recordID)
	return recordID
}

func (s *ConsoleUiService) GetCommandName() string {
	var commandName string
	fmt.Print(">")
	fmt.Scan(&commandName)
	return commandName
}

func (s *ConsoleUiService) OutputUser(user models.User) {
	fmt.Println("User Information:")
	fmt.Println("Username:", user.Username)
	fmt.Println("Password:", user.Password)
}

func (s *ConsoleUiService) OutputRecord(record models.Record) {
	fmt.Println("Record Information:")
	fmt.Println("Name:", record.Name)
	fmt.Println("Content:", record.Content)
	fmt.Println("Cost:", record.Cost)
}

func (s *ConsoleUiService) OutputString(str string) {
	fmt.Println(str)
}