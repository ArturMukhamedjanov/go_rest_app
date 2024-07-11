package commands

import(
	"go_rest_app/client/services"
)

type GetAllRecordsCommand struct{}

func  (c *GetAllRecordsCommand) Name() string{
	return "get_all_records"
}

func (c *GetAllRecordsCommand) Run(uiService service.UiService, serverService service.ServerService){
	records, err := serverService.GetAllRecords()
	if err != nil{
		uiService.OutputString("unable to send request")
		return
	}
	for _, record := range records{
		uiService.OutputRecord(record)
	}
}