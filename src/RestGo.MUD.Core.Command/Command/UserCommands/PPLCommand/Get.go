package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

type GetCommand struct{}

func (c *GetCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Get", false
}

func init() {
	UserCommands.RegisterCommand(&GetCommand{})
}

// put actionObj in conatinerObj
// put xxx in yyy
func (c *GetCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//拆解命令
	actionNum, actionObjName, conatinerNum, containerName := utility.ParserGetPutCommand(args)
	if actionObjName == "" {
		mudconn.SendMessage("你想拿什麼物品？")
		return
	}

	room, _ := RoomHelper.GetRoom(mudconn.User.RoomID) //假設房間一定存在
	//從房間內拿取
	if containerName == "" {
		actionObject, err := room.GetOut(actionNum, actionObjName)
		if err != nil {
			mudconn.SendMessage("這裏沒有那個東西。")
			return
		}

		if !actionObject.HaveCapability(BasicDefinition.CanBeMove) {
			mudconn.SendFMessage("你沒有辦法拿起%s。", actionObject.GetObjectBasic().Name_CH)
			return
		}

		mudconn.User.PutIn(actionObject)
		mudconn.SendFMessage("你撿起了%s。", actionObject.GetObjectBasic().Name_CH)
		return
	}

	//房間中有沒有該容器
	containerObj, err := room.ContainerPure.GetObjPoniter(conatinerNum, containerName)
	if err != nil {
		//房間中沒有該容器，試著從身上拿
		containerObj, err = mudconn.User.ContainerPure.GetObjPoniter(conatinerNum, containerName)
		if err != nil {
			mudconn.SendMessage(err.Error())
			return
		}
	}

	//確認是不是容器
	container, OK := containerObj.(*Container.ContainerObject)
	if !OK {
		mudconn.SendFMessage("你無法從%s當中拿出任何東西。", containerObj.GetObjectBasic().Name_CH)
		return
	}

	actionObject, err := container.GetOut(actionNum, actionObjName)
	if err != nil {
		mudconn.SendFMessage("%s裏沒有那個東西。", container.Name_CH)
		return
	}

	mudconn.User.PutIn(actionObject)
	mudconn.SendFMessage("你從%s拿出%s。", container.GetObjectBasic().Name_CH, actionObject.GetObjectBasic().Name_CH)

	return
}
