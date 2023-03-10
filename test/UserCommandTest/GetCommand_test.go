package UserCommandTest

import (
	"strings"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
)

// 將自身的物品放入背包裏
func TestGetCommand_Execute_Single(t *testing.T) {
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
	//設定房間內物品，一個背包，一把小刀，一座高山
	PutInBag(&room.ContainerPure)
	PutInKnife(&room.ContainerPure)
	PutInMontain(&room.ContainerPure)
	//背包裏放碎玻璃
	container := room.Items[0].(*Container.ContainerObject)
	PutInGlassess(&container.ContainerPure)

	//設定情境，地上沒有物品時
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"", "你想拿什麼物品？"},
		{"xxx", "這裏沒有那個東西。"},
		{"k", "你撿起了小刀。"},
		{"g from b", "你從背包拿出碎玻璃。"},
		{"g b", "背包裏沒有那個東西。"},
		{"b", "你撿起了背包。"},
		{"mon", "你沒有辦法拿起一座高山。"},
	}
	Command := &PPLCommand.GetCommand{}
	for _, tc := range testCases {
		ReceivedMessage = "" //每次都要先清空，避免受到上次的回傳影響
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
}

func TestGetCommand_Execute_AllFromRoom(t *testing.T) {
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
	//設定房間內物品，一個背包，一把小刀，一座高山
	PutInBag(&room.ContainerPure)
	PutInMontain(&room.ContainerPure)
	//背包裏放碎玻璃
	container := room.Items[0].(*Container.ContainerObject)
	PutInGlassess(&container.ContainerPure)

	//設定情境，地上沒有物品時
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"all from b", "你從背包拿出碎玻璃。"},
		{"all b", "背包裏空空如也。"},
		{"all xxx", "這裏沒有那個東西。"},
		{"b", "你撿起了背包。"},
		{"all", "你沒有辦法拿起一座高山。"},
	}
	Command := &PPLCommand.GetCommand{}
	for _, tc := range testCases {
		ReceivedMessage = "" //每次都要先清空，避免受到上次的回傳影響
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
}

func TestGetCommand_Execute_CanNotGet(t *testing.T) {
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
	//設定房間內物品，兩個金庫，第二個金庫是空的
	PutInSafeBox(&room.ContainerPure)
	PutInSafeBox(&room.ContainerPure)
	//公車裏放山
	container := room.Items[0].(*Container.ContainerObject)
	PutInMontain(&container.ContainerPure)

	//身上放小刀
	PutInKnife(&mudconn.User.ContainerPure)

	//設定情境
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"all from s", "你沒有辦法從金庫裏拿出一座高山。"},
		{"s", "你沒有辦法拿起金庫。"},
		{"all xxx", "這裏沒有那個東西。"},
		{"xxx s", "金庫裏沒有那個東西。"},
		{"all 2.s", "金庫裏空空如也。"},
		{"xxx from k", "你無法從小刀當中拿出任何東西。"},
	}
	Command := &PPLCommand.GetCommand{}
	for _, tc := range testCases {
		ReceivedMessage = "" //每次都要先清空，避免受到上次的回傳影響
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}
}
