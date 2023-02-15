package PPLCommand

import (
	"fmt"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

type InventoryCommand struct{}

func (c *InventoryCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Inventory", true
}

func init() {
	UserCommands.RegisterCommand(&InventoryCommand{})
}

func (c *InventoryCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	msg := "目前攜帶有:"
	for _, v := range mudconn.User.Items {
		msg += fmt.Sprintf("\n%s (%s)", v.GetObjectBasic().Name_CH, v.GetObjectBasic().Name_EN)
	}
	mudconn.SendMessage(msg)
	return
}
