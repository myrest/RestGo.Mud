package UserCommandTest

import (
	"fmt"
	"strings"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"

	//註冊命令用
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
	//初始化物件
	//_ "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

func TestTestCommand_Execute(t *testing.T) {
	// 建立一個模擬的MudClient物件
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User: Player.User{
			Password: "test_password",
		},
	}
	receivedMessage := ""

	// 建立一個模擬的SendFMessage
	StructCollection.SendFMessage = func(s *StructCollection.MudClient, msg string, args ...interface{}) {
		receivedMessage = fmt.Sprintf(msg+"\n", args...)
	}

	// 執行Execute方法
	Command := &PPLCommand.TestCommand{}
	Command.Execute("test", mudconn)

	// 檢查Execute方法是否發送了正確的訊息
	receivedMessage = strings.TrimSpace(receivedMessage)
	expectedMessage := "你憑空造出了一個背包。"
	if receivedMessage != expectedMessage {
		t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", expectedMessage, receivedMessage)
	}
}
