package main

import (
	"go_rest_app/client/app"
	service "go_rest_app/client/services"
)

func main() {
	app := app.NewApp(&service.ConsoleUiService{}, service.NewRestServerService())
	app.Start()
}
