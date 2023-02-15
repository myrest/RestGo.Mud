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

func TestPutCommand_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	PutInContainer("BagGUID-123456", &user.ContainerPure)
	PutInKnife("KnifeGUID-123456", &user.ContainerPure)
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
		{"xxxx", "你想放什麼物品進哪個物品？"},
		{"xxx bag", "你身上沒有那個東西。"},
		{"xxx knife", "它沒辦法放東西進去。"},
		//{"xxx xxx", "你身上沒有那個東西。"},
	}
	Command := &PPLCommand.PutCommand{}
	for _, tc := range testCases {
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
}
