package LoginCommands

import (
	"errors"
	"fmt"

	"github.com/logrusorgru/aurora/v4"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/CacheService"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
)

// Command is an interface for a command
type Command interface {
	Execute(string, *StructCollection.MudClient, *CacheService.LoginUser) (exitLoop bool, err error)
}

// CommandFunc is a function that creates a new Command
type CommandFunc func(string) Command

var commandFactories = make(map[int]CommandFunc)

func RegisterCommand(step int, input string, cmd Command) {
	if _, ok := commandFactories[step]; ok {
		panic(fmt.Sprintf("命令重覆定義[%s]", input))
	} else {
		commandFactories[step] = func(args string) Command { return cmd }
	}
}

// ProcessMessage processes a message from a telnet connection
func ProcessMessage(msg string, mudconn *StructCollection.MudClient) (exitLoop bool) {
	//先取得客戶的step在哪
	var cache = CacheService.CacheGetUser(mudconn.ConnectionID)
	cmdFunc, ok := commandFactories[cache.Step]
	if !ok {
		//未記錄，回到預設0
		fmt.Printf("步驟未記錄：%d", cache.Step)
		cache.SetStep(0)
		//重新取得commnad function
		cmdFunc, ok = commandFactories[cache.Step]
		if !ok {
			panic("登入預設步驟 0 未實作。")
		}
	}
	cmd := cmdFunc(msg)
	exitLoop, err := cmd.Execute(msg, mudconn, cache)
	if err != nil {
		mudconn.SendFMessage(err.Error())
	}
	return exitLoop
}

// //////login commands////////////////////////////////////////////////////////
type GetUserPasswordCommand struct{} //0
func (c *GetUserPasswordCommand) Execute(pwd string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, err error) {
	if pwd != cache.User.Password {
		if cache.LoginFailed < 4 {
			cache.SetLoginFailed()
			err = errors.New("密碼錯誤。請再輸入一次你的密碼：")
		} else {
			cache.RemoveUser()
			mudconn.SendFMessage("輸入錯誤密碼超過%d次，請重新登入。", aurora.Cyan(5))
			mudconn.Conn.Close()
		}
	} else {
		exitLoop = true
		mudconn.SendFMessage("登入成功。")
	}
	return
}

func init() {
	RegisterCommand(0, "", &GetUserPasswordCommand{})
}
