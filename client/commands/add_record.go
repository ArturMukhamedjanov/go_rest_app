package commands

import (
	service "go_rest_app/client/services"
)

type AddRecordCommand struct{}

func (c *AddRecordCommand) Name() string {
	return "add_record"
}

func (c *AddRecordCommand) Run(uiService service.UiService, serverService service.ServerService) {
	record := uiService.GetRecord()
	result := serverService.AddRecord(record)
	uiService.OutputString(result)
}