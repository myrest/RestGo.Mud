package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

type PutCommand struct{}

func (c *PutCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Put", false
}

func init() {
	UserCommands.RegisterCommand(&PutCommand{})
}

// put actionObj in conatinerObj
// put xxx in yyy
func (c *PutCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//拆解命令
	actionNum, actionObjName, conatinerNum, containerName := utility.ParserGetPutCommand(args)
	if actionObjName == "" || containerName == "" {
		mudconn.SendMessage("你想放什麼物品進哪個物品？")
		return
	}

	room, _ := RoomHelper.GetRoom(mudconn.User.RoomID) //假設房間一定存在
	PutToRoom := false
	//先試著取出身上conatinerObj的Pointer
	containerObj, err := mudconn.User.ContainerPure.GetObjPoniter(conatinerNum, containerName)
	if err != nil {
		//如果身上的取不到，試著取房間內的
		containerObj, err = room.ContainerPure.GetObjPoniter(conatinerNum, containerName)
		if err != nil {
			mudconn.SendMessage("這裏沒有那個東西。")
			return
		}
		PutToRoom = true
	}

	//先確認Container可以放東西
	container, OK := containerObj.(*Container.ContainerObject)
	if !OK {
		mudconn.SendFMessage("你無法將東西放入%s裏。", containerObj.GetObjectBasic().Name_CH)
		return
	}

	//取出要放入的物件
	actionObj, err := mudconn.User.ContainerPure.GetOut(actionNum, actionObjName)
	if err != nil {
		mudconn.SendMessage("你身上沒有那個東西。")
		return
	}

	if containerObj.GetObjectBasic().ID == actionObj.GetObjectBasic().ID {
		mudconn.SendMessage("你試著看看能不能把你自己給裝進自己？")
		mudconn.User.ContainerPure.PutIn(actionObj)
		return
	}

	container.PutIn(actionObj)
	if PutToRoom {
		mudconn.SendFMessage("你把%s放入眼前的%s。", actionObj.GetObjectBasic().Name_CH, containerObj.GetObjectBasic().Name_CH)
	} else {
		mudconn.SendFMessage("你把%s放入身上的%s。", actionObj.GetObjectBasic().Name_CH, containerObj.GetObjectBasic().Name_CH)
	}

	return
}
