package commands

import service "go_rest_app/client/services"

type HelpCommand struct{}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Run(uiService service.UiService, serverService service.ServerService) {
	uiService.OutputString("List of commands:")
	uiService.OutputList(GetCommandsList())
}