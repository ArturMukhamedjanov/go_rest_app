package commands

import(
	"go_rest_app/client/services"
)

type Command interface {
	Name() string
	Run(service.UiService, service.RestServerService)
}

var commands = make(map[string]Command)

func CommandRegistration(cmd Command) {
	commands[cmd.Name()] = cmd
}

func GetCommandAndExecute(name string, uiService service.UiService, serverService service.RestServerService) {
	commands[name].Run(uiService, serverService)
}
