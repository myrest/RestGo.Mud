package PPLCommand

import (
	"strings"

	"google.golang.org/grpc/status"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
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
	room, _ := RoomHelper.GetRoom(mudconn.User.RoomID)
	if actionObjName == "" {
		mudconn.SendMessage("你想丟掉什麼物品？")
		return
	}

	if actionNum == 1 && strings.ToLower(actionObjName) == "all" {
		c.DropAll(mudconn, room)
		return
	}

	obj, err := mudconn.User.GetOut(actionNum, actionObjName)
	if err != nil {
		mudconn.SendMessage("你沒有那個東西。")
		return
	}

	room.PutIn(obj)
	mudconn.SendFMessage("你丟掉%s。", obj.GetObjectBasic().Name_CH)

	return
}

func (*DropCommand) DropAll(mudconn *StructCollection.MudClient, room *Room.Room) {
	if len(mudconn.User.Items) > 0 {
		var itemsName []string
		for _, v := range mudconn.User.Items {
			itemsName = append(itemsName, v.GetObjectBasic().Name_EN)
		}
		for _, itemName := range itemsName {
			obj, err := mudconn.User.GetOut(1, itemName)
			if err != nil {
				switch status.Code(err) {
				case BasicDefinition.ObjectCannotMove:
					mudconn.SendFMessage("你沒有辦法丟掉%s", obj.GetObjectBasic().Name_CH)
				default:
					mudconn.SendMessage("你身上沒有那個東西。")
				}
			} else {
				room.PutIn(obj)
				mudconn.SendFMessage("你丟掉%s。", obj.GetObjectBasic().Name_CH)
			}

		}
	} else {
		mudconn.SendMessage("你身上空空如也。")
	}
}
