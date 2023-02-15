package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

type LookCommand struct{}

func (c *LookCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Look", true
}

func init() {
	UserCommands.RegisterCommand(&LookCommand{})
}

func (c *LookCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//取出房間
	r, err := RoomHelper.GetRoom(mudconn.User.RoomID)
	if err != nil {
		mudconn.SendMessage(err.Error())
		return
	}

	//沒加任何參數，表示看的是房間
	if args == "" {
		mudconn.SendMessage(c.LookRoom(r))
		return
	}

	//有參數，表示看的是物品
	num, objName := utility.ParserObjectNumber(args)
	//先查環境地上有沒有該物品
	obj, err := r.GetObjPoniter(num, objName)
	if err == nil {
		mudconn.SendMessage(obj.GetObjectBasic().Description_Ground)
		return
	}

	//找身上物品
	obj, err = mudconn.User.ContainerPure.GetObjPoniter(num, objName)
	if err == nil {
		mudconn.SendMessage(obj.GetObjectBasic().Description_Look)
		return
	}

	mudconn.SendMessage("你找不到那樣物品。")
	return
}
