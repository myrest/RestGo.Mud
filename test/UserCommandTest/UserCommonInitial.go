package UserCommandTest

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ExitsPosition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"
)

var ReceivedMessage string

func Begin() {
	// 建立一個模擬的SendFMessage
	StructCollection.SendFMessage = func(s *StructCollection.MudClient, msg string, args ...interface{}) {
		ReceivedMessage = fmt.Sprintf(msg+"\n", args...)
	}

	// 建立一個模擬的SendMessage
	StructCollection.SendMessage = func(s *StructCollection.MudClient, msg ...interface{}) {
		ReceivedMessage = fmt.Sprintln(msg...)
	}

}

func InitialUser() Player.User {
	rtn := Player.User{
		Creature: Creature.Creature{
			ID:               uuid.NewV4().String(),
			Name:             "TestUser001_Name",
			Description:      "這個是一個用來測試的Lv1帳號",
			Title:            "測試Lv1專用",
			AttributeBasic:   Creature.UserAttributes{},
			AttributeCurrent: Creature.UserAttributes{},
			Career:           Creature.Thief,
			Race:             Creature.Human,
			RoomID:           1,
			Level:            1,
			ContainerPure: Container.ContainerPure{
				Items: []BasicDefinition.IObjectBasic{},
			},
		},
		Password: "test_password",
		Gender:   Player.Male,
		Settings: Player.Configuration{
			Prompt: "*測試用：生命%h% 精神%m% 體力%v%*",
		},
	}
	return rtn
}

func PutInBag(container *Container.ContainerPure) {
	bag := &Container.ContainerObject{
		ObjectBasic: BasicDefinition.ObjectBasic{
			ID:                   uuid.NewV4().String(),
			Name_EN:              "Bag",
			Name_CH:              "背包",
			Level:                1,
			Description_List:     "堪堪能用的背包",
			Description_Ground:   "一個被人放在地上的背包。",
			Description_Look:     "一個好用、簡便可以放物品的背包。",
			Weight:               10,
			Pricing:              100,
			Capability:           []BasicDefinition.BasicCapability{BasicDefinition.CanBeMove},
			ObjectType:           1,
			DestroyWhenZeroQuota: false,
			AllowExecuteTimes:    0,
			Decoration:           []string{},
		},
		ContainerPure: Container.ContainerPure{
			Items: []BasicDefinition.IObjectBasic{},
		},
	}
	container.PutIn(bag)
}

func PutInKnife(container *Container.ContainerPure) {
	object := &BasicDefinition.ObjectBasic{
		ID:                 uuid.NewV4().String(),
		Name_EN:            "Knife",
		Name_CH:            "小刀",
		Level:              1,
		Description_List:   "鈍掉的小刀",
		Description_Ground: "這是一把放在地上的小刀",
		Description_Look:   "一把沒什麼用的小刀，看來不怎麼樣。",
		Weight:             10,
		Pricing:            100,
		Capability:         []BasicDefinition.BasicCapability{BasicDefinition.CanBeMove},
	}
	container.PutIn(object)
}

func PutInHelmet(container *Container.ContainerPure) {
	object := &BasicDefinition.ObjectBasic{
		ID:                 uuid.NewV4().String(),
		Name_EN:            "Helmet",
		Name_CH:            "安全帽",
		Level:              1,
		Description_List:   "破掉的安全帽",
		Description_Ground: "這是一頂破掉被丟棄的安全帽",
		Description_Look:   "破掉的安全帽，有總比沒有的好",
		Weight:             10,
		Pricing:            100,
		Capability:         []BasicDefinition.BasicCapability{BasicDefinition.CanBeMove},
	}
	container.PutIn(object)
}

func PutInMontain(container *Container.ContainerPure) {
	object := &BasicDefinition.ObjectBasic{
		ID:                 uuid.NewV4().String(),
		Name_EN:            "Montain",
		Name_CH:            "一座高山",
		Level:              1,
		Description_Ground: "這是一座很高的高山",
		Description_Look:   "高山，只能遠觀啊～～",
		Weight:             10,
		Pricing:            100,
	}
	container.PutIn(object)
}

func PutInGlassess(container *Container.ContainerPure) {
	object := &BasicDefinition.ObjectBasic{
		ID:                 uuid.NewV4().String(),
		Name_EN:            "Glasse",
		Name_CH:            "碎玻璃",
		Level:              1,
		Description_List:   "尖銳的碎玻璃",
		Description_Ground: "一片閃閃發亮的碎玻璃掉在地上。",
		Description_Look:   "閃閃發亮的碎玻璃，看起來真是耀眼。",
		Weight:             10,
		Pricing:            100,
	}
	container.PutIn(object)
}

func InitialRoom() *Room.Room {
	rtn := &Room.Room{
		ContainerPure: Container.ContainerPure{
			Items: []BasicDefinition.IObjectBasic{},
		},
		Title:       "測試房間001",
		Ceiling:     Room.Light,
		Description: "這是測試房間001Description",
		Geography: StructCollection.WorldGeography{
			World:  "亞爾斯隆世界",
			Region: "區域次中",
			Area:   "區域最小",
			RoomID: 1,
		},
		Exits: map[ExitsPosition.ExitName]ExitsPosition.Exit{},
	}
	RoomHelper.AddRoom(rtn)
	return rtn
}
