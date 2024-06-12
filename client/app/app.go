package app

import (
	"go_rest_app/client/commands"
	"go_rest_app/client/services"
	"os"
)

type App struct {
	uiService     service.UiService
	serverService service.RestServerService
}

func NewApp(uS service.UiService, sS service.RestServerService) App {
	return App{
		uiService:     uS,
		serverService: sS,
	}
}

func (app *App) Start(){
	app.loadBasicCommands()
	app.startListening()
}

func (app *App) loadBasicCommands(){
	commands.CommandRegistration(&commands.RegisterCommand{})
	
}

func (app *App) startListening() {
	for {
		commandName := app.uiService.GetCommandName()
		if commandName == "exit" {
			app.uiService.OutputString("Exiting programm...")
			os.Exit(0)
		}
		commands.GetCommandAndExecute(commandName, app.uiService, app.serverService)
	}
}
