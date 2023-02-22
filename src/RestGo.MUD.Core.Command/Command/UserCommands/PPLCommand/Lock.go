package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

type LockCommand struct{}

func (c *LockCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Lock", false
}

func (c *LockCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	mudconn.SendMessage("This is lock command")
	return
}

func init() {
	UserCommands.RegisterCommand(&LockCommand{})
}
