package PPLCommand

import (
	"strings"

	"google.golang.org/grpc/status"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
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

// get all
// get xxx
// get xxx from yyy
// get xxx yyy
// get all xxx
// get all from yyy
func (c *GetCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//拆解命令
	actionNum, actionObjName, conatinerNum, containerName := utility.ParserGetPutCommand(args)
	if actionObjName == "" {
		mudconn.SendMessage("你想拿什麼物品？")
		return
	}
	room, _ := RoomHelper.GetRoom(mudconn.User.RoomID) //假設房間一定存在

	//拿房間--全部
	if actionNum == 1 && strings.ToLower(actionObjName) == "all" && containerName == "" {
		c.GetAllFromRoom(room, mudconn)
		return
	}

	//拿Container--全部
	if actionNum == 1 && strings.ToLower(actionObjName) == "all" && containerName != "" {
		c.GetAllFromContainer(conatinerNum, containerName, mudconn, room)
		return
	}

	//從房間內拿取物品
	if containerName == "" {
		c.GetFromRoom(actionNum, actionObjName, mudconn, room)
		return
	}

	//從Container裏拿取物品
	//房間中有沒有該容器
	containerObj, err := room.GetObjPoniter(conatinerNum, containerName)
	if err != nil {
		//房間中沒有該容器，試著從身上拿
		containerObj, err = mudconn.User.GetObjPoniter(conatinerNum, containerName)
		if err != nil {
			mudconn.SendMessage("這裏沒有那個東西。")
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
		switch status.Code(err) {
		case BasicDefinition.ObjectCannotMove:
			mudconn.SendFMessage("你沒有辦法從%s裏拿出%s", actionObject.GetObjectBasic().Name_CH)
		default:
			mudconn.SendFMessage("%s裏沒有那個東西。", container.GetObjectBasic().Name_CH)
		}
	} else {
		mudconn.User.PutIn(actionObject)
		mudconn.SendFMessage("你從%s拿出%s。", container.GetObjectBasic().Name_CH, actionObject.GetObjectBasic().Name_CH)
	}
	return
}

func (*GetCommand) GetAllFromRoom(room *Room.Room, mudconn *StructCollection.MudClient) {
	if len(room.Items) == 0 {
		mudconn.SendMessage("這個房間空空如也。")
		return
	}
	var itemNames []string
	for _, item := range room.Items {
		itemNames = append(itemNames, item.GetObjectBasic().Name_EN)
	}
	for _, itemName := range itemNames {
		item, err := room.GetOut(1, itemName)
		if err != nil {
			switch status.Code(err) {
			case BasicDefinition.ObjectCannotMove:
				mudconn.SendFMessage("你沒有辦法拿起%s。", item.GetObjectBasic().Name_CH)
			default:
				mudconn.SendMessage("這裏沒有那個東西。")
			}
		} else {
			mudconn.User.PutIn(item)
			mudconn.SendFMessage("你撿起了%s。", item.GetObjectBasic().Name_CH)
		}
	}
}

func (*GetCommand) GetAllFromContainer(conatinerNum int, containerName string, mudconn *StructCollection.MudClient, room *Room.Room) {
	containerObj, err := room.GetObjPoniter(conatinerNum, containerName)
	if err != nil {
		//房間中沒有該容器，試著從身上拿
		containerObj, err = mudconn.User.GetObjPoniter(conatinerNum, containerName)
		if err != nil {
			mudconn.SendMessage("這裏沒有那個東西。")
			return
		}
	}

	//確認是不是容器
	container, OK := containerObj.(*Container.ContainerObject)
	if !OK {
		mudconn.SendFMessage("你無法從%s當中拿出任何東西。", containerObj.GetObjectBasic().Name_CH)
		return
	}

	if len(container.Items) < 1 {
		mudconn.SendFMessage("%s裏空空如也。", container.GetObjectBasic().Name_CH)
		return
	}

	var itemNames []string
	for _, item := range container.Items {
		itemNames = append(itemNames, item.GetObjectBasic().Name_EN)
	}
	for _, itemName := range itemNames {
		item, err := container.GetOut(1, itemName)
		if err != nil {
			switch status.Code(err) {
			case BasicDefinition.ObjectCannotMove:
				mudconn.SendFMessage("你沒有辦法從%s裏拿出%s。", container.GetObjectBasic().Name_CH, item.GetObjectBasic().Name_CH)
			default:
				mudconn.SendFMessage("%s裏沒有那個東西。", container.GetObjectBasic().Name_CH)
			}
		} else {
			mudconn.User.PutIn(item)
			mudconn.SendFMessage("你從%s拿出%s。", container.GetObjectBasic().Name_CH, item.GetObjectBasic().Name_CH)
		}
	}
}

func (*GetCommand) GetFromRoom(actionNum int, actionObjName string, mudconn *StructCollection.MudClient, room *Room.Room) {
	actionObject, err := room.GetOut(actionNum, actionObjName)
	if err != nil {
		switch status.Code(err) {
		case BasicDefinition.ObjectCannotMove:
			mudconn.SendFMessage("你沒有辦法拿起%s。", actionObject.GetObjectBasic().Name_CH)
		default:
			mudconn.SendMessage("這裏沒有那個東西。")
		}
	} else {
		mudconn.User.PutIn(actionObject)
		mudconn.SendFMessage("你撿起了%s。", actionObject.GetObjectBasic().Name_CH)
	}
}
