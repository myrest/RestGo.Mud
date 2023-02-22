package UserCommandTest

import (
	"strings"
	"testing"

	//註冊命令用
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	//初始化物件
	//_ "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

func TestTestCommand_Execute(t *testing.T) {
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
	InitialRoom()

	//設定情境，地上沒有該物品時
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"", "你想做啥？"},
		{"bag", "你憑空造出了一個袋子。"},
		{"Helmet", "你憑空造出了一個安全帽。"},
	}
	Command := &PPLCommand.TestCommand{}
	for _, tc := range testCases {
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
}
