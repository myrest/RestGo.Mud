package PPLCommand

import (
	uu "github.com/satori/go.uuid"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
)

type TestCommand struct{}

func (c *TestCommand) Settings() (Fullkey string, SupportFuzzyMatch bool) {
	return "Test", true
}

func (c *TestCommand) Execute(args string, mudconn *StructCollection.MudClient) (quit bool) {
	//測試物件
	bag := &Container.Bag{
		ContainerBasic: BasicDefinition.ContainerBasic{
			ObjectBasic: BasicDefinition.ObjectBasic{
				ID:                 uu.NewV4().String(),
				Name_EN:            "Bag",
				Name_CH:            "背包",
				Level:              1,
				Description_Ground: "一個放物品的背包",
				Description_Look:   "這是一個可以放物品的背包",
				Weight:             10,
				Pricing:            100,
			},
			ContainerPure: BasicDefinition.ContainerPure{
				Items: []BasicDefinition.IObjectBasic{},
			},
		},
	}
	mudconn.User.PutIn(bag)
	mudconn.SendFMessage("你憑空造出了一個%s。", bag.Name_CH)
	return
}

func init() {
	UserCommands.RegisterCommand(&TestCommand{})
}
