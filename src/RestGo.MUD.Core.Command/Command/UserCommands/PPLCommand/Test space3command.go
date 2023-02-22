package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

type TestSpace3CommandCommand struct{}

func (c *TestSpace3CommandCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Test Space 3Command", true
}

func (c *TestSpace3CommandCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	mudconn.SendFMessage("This is a test 3 space command")
	return
}

func init() {
	UserCommands.RegisterCommand(&TestSpace3CommandCommand{})
}
