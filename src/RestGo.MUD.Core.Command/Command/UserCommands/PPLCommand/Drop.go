package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

type DropCommand struct{}

func (c *DropCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Drop", false
}

func init() {
	UserCommands.RegisterCommand(&DropCommand{})
}

// drop xxx
func (c *DropCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//拆解命令
	actionNum, actionObjName, _, _ := utility.ParserGetPutCommand(args)
	if actionObjName == "" {
		mudconn.SendMessage("你想丟掉什麼物品？")
		return
	}

	obj, err := mudconn.User.GetOut(actionNum, actionObjName)
	if err != nil {
		mudconn.SendMessage("你沒有那個東西。")
		return
	}

	room, _ := RoomHelper.GetRoom(mudconn.User.RoomID)
	room.PutIn(obj)
	mudconn.SendFMessage("你丟掉%s。", obj.GetObjectBasic().Name_CH)

	return
}
