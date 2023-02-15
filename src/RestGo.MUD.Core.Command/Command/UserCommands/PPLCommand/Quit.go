package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

type QuitCommand struct{}

func (c *QuitCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Quit", false
}

func (c *QuitCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	mudconn.SendFMessage("\n\n\n\n\n\n\n\n穿越了黑暗深邃的時空之門, 突然眼前一亮...\n\n\n\n\n\n\n\n完蛋了! 又在電腦前面睡過頭了... 回想剛才種種, 彷彿經歷了一場夢幻....")
	return true
}

func init() {
	UserCommands.RegisterCommand(&QuitCommand{})
}
