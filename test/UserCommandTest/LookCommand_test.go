package UserCommandTest

import (
	"strings"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"

	//註冊命令用
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
	//初始化物件
	//_ "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

// todo:還沒寫完
func TestLookCommand_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	room := InitialRoom()
	PutInKnife("KnifeGUID-123456", &user.ContainerPure)
	PutInGlassess("GlassGUID-123456", &room.ContainerPure)
	PutInContainer("BagGUID-123456", &room.ContainerPure)
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User:         user,
	}

	//設定情境，地上沒有物品時
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"xxx", "你找不到那樣物品。"},
		{"g", "一片閃閃發亮的碎玻璃掉在地上。"},
		{"k", "一把沒什麼用的小刀，看來不怎麼樣。"},
		{"b", "一個被人放在地上的背包。"},
		//{"", ""},
	}
	Command := &PPLCommand.LookCommand{}
	for _, tc := range testCases {
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
}
