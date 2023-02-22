package UserCommands

import (
	"fmt"
	"strings"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

// Command is an interface for a command
type IUserCommand interface {
	Execute(string, *StructCollection.MudClient) (quit bool)
	Settings() (Fullkey string, DoNotSupportFuzzyMatch bool)
}

var commandFactories = make(map[string]IUserCommand)

func RegisterCommand(cmd IUserCommand) {
	fullKey, _ := cmd.Settings()
	lowerCmdKey := strings.ToLower(fullKey)
	if _, ok := commandFactories[lowerCmdKey]; ok {
		panic(fmt.Sprintf("命令重覆定義[%s]", fullKey))
	} else {
		commandFactories[lowerCmdKey] = cmd
	}
}

// ProcessMessage processes a message from a telnet connection
func ProcessMessage(msg string, mudconn *StructCollection.MudClient) (isQuit bool) {
	msg = strings.TrimSpace(msg)
	if len(msg) > 0 {
		// Parse the message and execute the corresponding command
		cmdName, cmdArgs := utility.SplitCommand(msg)
		cmdFunc := getCommand(cmdName)

		if cmdFunc == nil {
			fmt.Fprintln(mudconn.Conn, "Invalid command.")
			return
		}

		isQuit = cmdFunc.Execute(cmdArgs, mudconn)
	}
	//結束後要顯示Prompt
	if !isQuit {
		mudconn.SendMessage("\r\n" + mudconn.User.Prompt())
	}
	return
}

func getCommand(inputMessage string) IUserCommand {
	inputMessage = strings.ToLower(inputMessage)
	//去除單引號
	if len(inputMessage) >= 2 && inputMessage[0] == '\'' && inputMessage[len(inputMessage)-1] == '\'' {
		inputMessage = inputMessage[1 : len(inputMessage)-1]
	}

	for commandKey, userCommand := range commandFactories {
		//完全相同
		if commandKey == inputMessage {
			return userCommand
		}

		//開頭相同，且命令中沒有空格，且支援Fuzzy Search
		if _, SupportFuzzyMatch := userCommand.Settings(); strings.HasPrefix(commandKey, inputMessage) &&
			!strings.ContainsAny(commandKey, " \n\r") &&
			SupportFuzzyMatch {
			return userCommand
		}

		//如果Command key為兩字以上
		if strings.Contains(commandKey, " ") {
			//依空格柝分輸入的input及commandkey
			inputCommandWords := strings.Fields(inputMessage)
			commandKeyWords := strings.Fields(commandKey)
			match := true
			//只比對相同字數的Command
			if len(commandKeyWords) != len(inputCommandWords) {
				continue
			}
			for i := range commandKeyWords {
				if !strings.HasPrefix(commandKeyWords[i], inputCommandWords[i]) {
					match = false
					break
				}
			}
			if match {
				return userCommand
			}
		}
	}

	//以上都找不到就重新比對一次。如果是兩個單字以上相符cmd且找不到，就輸入多少，比對多少
	if strings.Contains(inputMessage, " ") {
		inputCommandWords := strings.Fields(inputMessage)
		for key, value := range commandFactories {
			keyWords := strings.Fields(key)
			match := true
			for i := range inputCommandWords {
				if !strings.HasPrefix(keyWords[i], inputCommandWords[i]) {
					match = false
					break
				}
			}
			if match {
				return value
			}
		}
	}
	return nil
}
