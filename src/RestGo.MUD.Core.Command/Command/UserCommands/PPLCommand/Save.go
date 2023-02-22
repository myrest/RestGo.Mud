package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Management/PlayerManagement"
)

type SaveCommand struct{}

func (c *SaveCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Save", false
}

func (c *SaveCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	if err := PlayerManagement.Save(mudconn.User); err != nil {
		mudconn.SendFMessage("存檔錯誤。\r\n" + err.Error())
	} else {
		mudconn.SendFMessage("存檔完成。")
	}
	return
}

func init() {
	UserCommands.RegisterCommand(&SaveCommand{})
}
