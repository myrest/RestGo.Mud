package Telnet

import (
	"bufio"
	"fmt"
	"time"

	"rest.com.tw/tinymud/src/RestGo.DataStorage/firebase"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/CreateNewUserCommand"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/LoginCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/Command/UserCommands"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/CacheService"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/World/WordInformation"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

// handleConnection handles a single telnet connection
func (s *TelnetServer) handleConnection(mudconn *StructCollection.MudClient) {
	if firebase.IsRunningInTestMode() {
		mudconn.SendFMessage("XXXXXXXXXXXX Local Mode XXXXXXXXXXXXXXXX")
	}

	// 超過10分鐘沒動作，中斷連線
	timeout := 10 * time.Minute
	timerStopConnection := time.AfterFunc(timeout, func() {
		defer s.closeClientConnection(mudconn)
	})

	defer s.closeClientConnection(mudconn)
	defer timerStopConnection.Stop()

	scanner := bufio.NewScanner(mudconn.Conn)
	welcomeMSG := utility.GetContentFile("Welcome.txt")
	//啟動時間，上線玩家，登入中玩家，總使用者
	mudconn.SendFMessage(welcomeMSG+"\r\n\r\n請輸入你的名字：",
		WordInformation.StartVirtualTime,
		WordInformation.StartTime,
		"B", "C", len(s.clients))

	exitLoop := false
	isCreated := false

	//建立帳號流程
	for {
		timerStopConnection.Reset(timeout)
		if !scanner.Scan() { //使用者中斷連後，要直接斷線
			return
		}
		msg := scanner.Text()
		exitLoop, isCreated = CreateNewUserCommand.ProcessMessage(msg, mudconn)
		if exitLoop {
			break
		}
	}

	if !isCreated {
		// 登入流程
		mudconn.SendFMessage("請輸入你的密碼：")
		for {
			timerStopConnection.Reset(timeout)
			if !scanner.Scan() { //使用者中斷連後，要直接斷線
				return
			}
			msg := scanner.Text()
			exitLoop := LoginCommands.ProcessMessage(msg, mudconn)
			if exitLoop {
				break
			}
		}
	}

	//登入或建立完成後，要從Cache移掉，放入MudClient物件裏。
	mudconn.User = CacheService.CacheGetUser(mudconn.ConnectionID).User
	CacheService.CacheGetUser(mudconn.ConnectionID).RemoveUser()

	mudconn.SendFMessage("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n突然螢幕變得黑暗, 有一股很強的力量把你吸進去....\n\n\n")
	mudconn.SendMessage(mudconn.User.Prompt())
	UserCommands.ProcessMessage("look", mudconn)

	// 開始正式進入遊戲
	for {
		timerStopConnection.Reset(timeout)
		if !scanner.Scan() { //使用者中斷連後，要直接斷線
			fmt.Println("使用者斷線")
			break
		}
		msg := scanner.Text()
		quit := UserCommands.ProcessMessage(msg, mudconn)
		if quit {
			break
		}
	}
}

func (s *TelnetServer) closeClientConnection(mudconn *StructCollection.MudClient) {
	CacheService.RemoveAllCacheByConnectionID(mudconn.ConnectionID)
	s.RemoveClient(mudconn)
	mudconn.Conn.Close()
}

func (s *TelnetServer) RemoveClient(elem *StructCollection.MudClient) {
	for i, v := range s.clients {
		if v == elem {
			s.clients[i] = s.clients[len(s.clients)-1]
			s.clients = s.clients[:len(s.clients)-1]
			break
		}
	}
}
