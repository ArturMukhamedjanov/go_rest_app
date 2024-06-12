package commands

import service "go_rest_app/client/services"

type RegisterCommand struct{}

func (c *RegisterCommand) Name() string {
	return "register"
}

func (c *RegisterCommand) Run(uiService service.UiService, serverService service.RestServerService) {
	user := uiService.GetUser()
	result := serverService.RegisterUser(user)
	uiService.OutputString(result)
}

