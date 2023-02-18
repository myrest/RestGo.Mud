package UserCommandTest

import (
	"strings"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
)

// todo:還沒寫完
func TestLookCommand_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	room := InitialRoom()
	PutInKnife(&user.ContainerPure)
	PutInHelmet(&room.ContainerPure)
	PutInBag(&room.ContainerPure)
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User:         user,
	}

	//設定情境，地上沒有該物品時
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"xxx", "你找不到那樣物品。"},
		{"k", "一把沒什麼用的小刀，看來不怎麼樣。"},
		{"b", "一個被人放在地上的背包。"},
		{"", "測試房間001\n明顯的出口有：\n這是測試房間001Description\n破掉的安全帽  (Helmet)\n堪堪能用的背包 (Bag)"},
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

func TestLookIntoObjCommand_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	room := InitialRoom()
	PutInBag(&room.ContainerPure)
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User:         user,
	}

	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"in", "你找不到那樣物品。"},
		{"into", "你找不到那樣物品。"},
		{"in xxx", "你找不到那樣物品。"},
		{"in b", "你看到裏面有："},
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

func TestLookMultipleObjCommand_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	room := InitialRoom()
	PutInGlassess(&room.ContainerPure)
	PutInGlassess(&room.ContainerPure)
	PutInBag(&room.ContainerPure)
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User:         user,
	}

	//設定情境，地上沒有該物品時
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"", "測試房間001\n明顯的出口有：\n這是測試房間001Description\n(2)尖銳的碎玻璃  (Glasse)\n    堪堪能用的背包 (Bag)"},
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
