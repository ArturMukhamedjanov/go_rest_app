package commands

import(
	"go_rest_app/client/services"
)

type GetRecordsByNameCommand struct{}

func  (c *GetRecordsByNameCommand) Name() string{
	return "get_records_by_name"
}

func (c *GetRecordsByNameCommand) Run(uiService service.UiService, serverService service.ServerService){
	name := uiService.GetCommandName()
	records, err := serverService.GetRecordsByName(name)
	if err != nil{
		uiService.OutputString("unable to send request")
		return
	}
	for _, record := range records{
		uiService.OutputRecord(record)
	}
}