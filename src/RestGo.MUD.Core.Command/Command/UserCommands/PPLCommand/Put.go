package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
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

	//取出身上conatinerObj的Pointer
	containerObj, err := mudconn.User.ContainerPure.GetObjPoniter(conatinerNum, containerName)
	if err != nil {
		mudconn.SendMessage(err.Error())
		return
	}

	/*
		//檢查obj是否能放東西
		_, ok := containerObj.(*BasicDefinition.ContainerBasic)
		if !ok {
			mudconn.SendMessage("它沒辦法放東西進去。")
			return
		}
	*/

	isContainer := false
	switch containerObj.(type) {
	case *BasicDefinition.ContainerBasic:
		isContainer = true
	}

	if !isContainer {
		mudconn.SendMessage("它沒辦法放東西進去。")
		return
	}

	//取出要放入的物件
	actionObj, err := mudconn.User.ContainerPure.GetOut(actionNum, actionObjName)
	if err != nil {
		mudconn.SendMessage(err.Error())
		return
	}

	if containerObj.GetObjectBasic().ID == actionObj.GetObjectBasic().ID {
		mudconn.SendMessage("你試著看看能不能把你自己給裝進自己？")
		mudconn.User.ContainerPure.PutIn(actionObj)
	} else {
		//因為之前IsContainer已經判斷過了，所以這裏不應該出錯。
		container := containerObj.(*BasicDefinition.ContainerBasic)
		container.PutIn(actionObj)
		mudconn.SendFMessage("你把%s放入%s", actionObj.GetObjectBasic().Name_CH, containerObj.GetObjectBasic().Name_CH)
	}

	return
}
