package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

type Test2Space3CommandCommand struct{}

func (c *Test2Space3CommandCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Test 2Space 3Command", true
}

func (c *Test2Space3CommandCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	mudconn.SendFMessage("This is a test 2space command")
	return
}

func init() {
	UserCommands.RegisterCommand(&Test2Space3CommandCommand{})
}
