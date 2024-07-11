package commands

import(
	"go_rest_app/client/services"
)

type GetRecordByIdCommand struct{}

func  (c *GetRecordByIdCommand) Name() string{
	return "get_record_by_id"
}

func (c *GetRecordByIdCommand) Run(uiService service.UiService, serverService service.ServerService){
	id := uiService.GetRecordID()
	record, err := serverService.GetRecordByID(id)
	if err != nil{
		uiService.OutputString("unable to send request")
		return
	}
	uiService.OutputRecord(*record)
}