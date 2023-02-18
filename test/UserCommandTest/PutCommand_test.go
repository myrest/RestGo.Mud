package UserCommandTest

import (
	"strings"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"

	//註冊命令用
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands/PPLCommand"
	//初始化物件
	//_ "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

// 將自身的物品放入背包裏
func TestPutCommandUserToSelf_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	PutInBag(&user.ContainerPure)
	PutInKnife(&user.ContainerPure)
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
		{"xxx knife", "你無法將東西放入小刀裏。"},
		{"xxx in knife", "你無法將東西放入小刀裏。"},
		{"xxx into knife", "你無法將東西放入小刀裏。"},
		{"xxx bag", "你身上沒有那個東西。"},
		{"b b", "你試著看看能不能把你自己給裝進自己？"},
		{"k b", "你把小刀放入袋子。"},
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

// 將自身的物品放入背包裏，檢查數量是否正確
func TestPutCommandUserToSelfCheckNum_Execute(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	PutInBag(&user.ContainerPure)
	PutInKnife(&user.ContainerPure)
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
		{"k b", "你把小刀放入背包。"},
	}
	Command := &PPLCommand.PutCommand{}
	for _, tc := range testCases {
		Command.Execute(tc.msg, mudconn)
		ReceivedMessage = strings.TrimSpace(ReceivedMessage)
		if ReceivedMessage != tc.expectedMessage {
			t.Errorf("Execute method does not send the correct message.\nExpected: %s\nActual: %s\n", tc.expectedMessage, ReceivedMessage)
		}
	}

	if len(mudconn.User.Items) != 1 {
		t.Errorf("玩家身上物品沒有減少。")
	}
	container := mudconn.User.Items[0].(*Container.ContainerObject)
	if len(container.Items) != 1 {
		t.Errorf("背包裏的物品沒有增加。")
	}
}

func TestPutFromUserToObjectInTheGround(t *testing.T) {
	//先Inject test需要的函式
	Begin()
	// 建立一個模擬的MudClient物件
	user := InitialUser()
	mudconn := &StructCollection.MudClient{
		Conn:         nil,
		ConnectionID: "123456",
		User:         user,
	}
	// 設定使用者身上物品，身上有兩把Knife
	PutInKnife(&mudconn.User.ContainerPure)
	PutInKnife(&mudconn.User.ContainerPure)

	//建立一個房間
	room := InitialRoom()
	//設定房間內物品
	PutInBag(&room.ContainerPure) //在地上放一個背包

	//設定情境，將身上物品放入地上背包
	testCases := []struct {
		msg             string
		expectedMessage string
	}{
		{"xxx bag", "你身上沒有那個東西。"},
		{"k b", "你把小刀放入眼前的背包。"},
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
