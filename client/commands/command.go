package commands

import (
	"go_rest_app/client/services"
)

type Command interface {
	Name() string
	Run(service.UiService, service.ServerService)
}

var commands = make(map[string]Command)
func CommandRegistration(cmd Command) {
	commands[cmd.Name()] = cmd
}

func GetCommandAndExecute(name string, uiService service.UiService, serverService service.ServerService) {
	commands[name].Run(uiService, serverService)
}

func GetCommandsList() []string {
	var commandList []string
	for name := range commands {
		commandList = append(commandList, name)
	}
	return commandList
}
