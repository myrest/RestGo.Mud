package PPLCommand

import (
	"strings"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
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

	//先檢查是不是Look In
	msgArr := strings.SplitN(strings.ToLower(args), " ", 2)
	if len(msgArr) == 2 && (msgArr[0] == "in" || msgArr[0] == "into") {
		num, objName := utility.ParserObjectNumber(msgArr[1])
		//先查環境地上有沒有該物品
		obj, err := r.GetObjPoniter(num, objName)
		if err == nil {
			if obj.GetObjectBasic().ObjectType == BasicDefinition.Container {
				container := obj.(*Container.ContainerObject)
				msg := "你看到裏面有：\n"
				msg = msg + container.ItemListForDisplay()
				mudconn.SendMessage(msg)
				return
			}
			mudconn.SendMessage(obj.GetObjectBasic().Name_CH, "需要判斷是不是可以裝水的容器。")
			return
		}
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
