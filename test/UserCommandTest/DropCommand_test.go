package UserCommandTest

import (
	"strings"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"

	//註冊命令用
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
)

func TestDropCommand_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User:         user,
	}

	//建立一個房間
	room := InitialRoom()
	//設定使用者身上物品
	PutInBag(&mudconn.User.ContainerPure)
	PutInKnife(&mudconn.User.ContainerPure)

	//設定情境
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"", "你想丟掉什麼物品？"},
		{"xxx", "你沒有那個東西。"},
		{"b", "你丟掉背包。"},
	}
	Command := &PPLCommand.DropCommand{}
	for _, tc := range testCases {
		ReceivedMessage = "" //每次都要先清空，避免受到上次的回傳影響
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
	if len(room.Items) < 1 {
		t.Errorf("物品沒有放到Room裏")
	}
}
