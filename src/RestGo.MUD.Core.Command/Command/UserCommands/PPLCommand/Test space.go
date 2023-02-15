package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

type TestSpaceCommand struct{}

func (c *TestSpaceCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Test Space", true
}

func (c *TestSpaceCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	mudconn.SendFMessage("This is a test space command")
	return
}

func init() {
	UserCommands.RegisterCommand(&TestSpaceCommand{})
}
