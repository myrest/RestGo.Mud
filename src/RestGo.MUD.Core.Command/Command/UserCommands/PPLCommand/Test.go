package PPLCommand

import (
	"strings"

	uuid "github.com/satori/go.uuid"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

type TestCommand struct{}

func (c *TestCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Test", true
}

func (c *TestCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//取出房間
	r, err := RoomHelper.GetRoom(mudconn.User.RoomID)
	if err != nil {
		mudconn.SendMessage(err.Error())
		return
	}
	var obj BasicDefinition.IObjectBasic
	switch strings.ToLower(args) {
	case "bag":
		obj = GetBag()
	case "helmet":
		obj = GetHelmet()
	default:
		mudconn.SendMessage("你想做啥？")
		return
	}

	r.PutIn(obj)
	mudconn.SendFMessage("你憑空造出了一個%s。", obj.GetObjectBasic().Name_CH)
	return
}

func init() {
	UserCommands.RegisterCommand(&TestCommand{})
}

func GetHelmet() BasicDefinition.IObjectBasic {
	objectBasic := BasicDefinition.ObjectBasic{
		ID:                 uuid.NewV4().String(),
		Name_EN:            "Helmet",
		Name_CH:            "安全帽",
		Level:              1,
		Description_List:   "破掉的安全帽",
		Description_Ground: "這是一頂破掉被丟棄的安全帽",
		Description_Look:   "破掉的安全帽，有總比沒有的好",
		ObjectType:         BasicDefinition.Basic,
		Capability:         []BasicDefinition.BasicCapability{BasicDefinition.CanBeMove},
	}
	return &objectBasic
}

func GetBag() BasicDefinition.IObjectBasic {
	objectBasic := BasicDefinition.ObjectBasic{
		ID:                 uuid.NewV4().String(),
		Name_EN:            "Bag",
		Name_CH:            "袋子",
		Level:              1,
		Description_List:   "堪堪能用的背包",
		Description_Ground: "可以用來裝些小物品的袋子",
		Description_Look:   "看起來破舊不堪，不太耐用的袋子。",
		ObjectType:         BasicDefinition.Container,
		Capability:         []BasicDefinition.BasicCapability{BasicDefinition.CanBeMove},
	}
	rtn := Container.ContainerObject{
		ObjectBasic: objectBasic,
	}
	return &rtn
}
