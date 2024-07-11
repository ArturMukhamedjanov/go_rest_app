package app

import (
	"go_rest_app/client/commands"
	"go_rest_app/client/services"
	"os"
)

type App struct {
	uiService     service.UiService
	serverService service.ServerService
}

func NewApp(uS service.UiService, sS service.ServerService) App {
	return App{
		uiService:     uS,
		serverService: sS,
	}
}

func (app *App) enterApp() {
	entered := false
	for !entered {
		app.uiService.OutputString("Type \"login\" to login in App. Type \"register\" to register in app")
		mode := app.uiService.GetCommandName()
		if mode == "register" {
			user := app.uiService.GetUser()
			result, id := app.serverService.RegisterUser(user)
			app.uiService.OutputString(result)
			if id != -1 {
				entered = true
				app.uiService.OutputString("Successfully entered app as " + user.Username)
				service.SetActiveUserId(id)
			}
		} else if mode == "login" {
			user := app.uiService.GetUser()
			id := app.serverService.AuthentificateUser(user)
			if id == -1 {
				app.uiService.OutputString("Wrond user data, try again")
			} else {
				entered = true
				app.uiService.OutputString("Successfully entered app as " + user.Username)
				service.SetActiveUserId(id)
			}
		} else {
			app.uiService.OutputString("Unknown mode. Try Again")
		}
	}
}

func (app *App) Start() {
	app.loadBasicCommands()
	app.enterApp()
	app.startListening()

}

func (app *App) loadBasicCommands() {
	commands.CommandRegistration(&commands.HelpCommand{})
	commands.CommandRegistration(&commands.AddRecordCommand{})
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
