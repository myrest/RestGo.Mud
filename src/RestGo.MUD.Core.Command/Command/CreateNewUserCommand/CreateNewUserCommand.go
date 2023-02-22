package CreateNewUserCommand

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/logrusorgru/aurora/v4"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/CacheService"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Management/PlayerManagement"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

// Command is an interface for a command
type Command interface {
	Execute(string, *StructCollection.MudClient, *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error)
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
func ProcessMessage(msg string, mudconn *StructCollection.MudClient) (exitLoop bool, isCreated bool) {
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
	exitLoop, isCreated, err := cmd.Execute(msg, mudconn, cache)
	if err != nil {
		mudconn.SendFMessage(err.Error())
	}
	return exitLoop, isCreated
}

// //////Create new user commands////////////////////////////////////////////////////////
type IsNewCreationCommand struct{} //0
func (c *IsNewCreationCommand) Execute(args string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	if err = utility.ValidateStringLengRange(args, 4, 20); err != nil {
		if len(args) == 0 {
			err = errors.New("請輸入你的名字：")
		}
		return
	}
	cache.User.ID = args
	pUser, _ := PlayerManagement.GetByID(args)
	if pUser != nil && pUser.ID != "" {
		//進入登入流程
		cache.User = *pUser
		cache.SetStep(0)
		exitLoop = true
	} else {
		mudconn.SendFMessage(aurora.Sprintf("這個名字聽起來很陌生，要使用%s做為你的名字嗎[%s/%s]？", aurora.Cyan(args), aurora.Cyan("y"), aurora.Cyan("N")))
		cache.SetStep(1)
	}
	return
}

type ConfirmCreationCommand struct{} //1
func (c *ConfirmCreationCommand) Execute(args string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	args = strings.ToLower(args)
	if args == "n" || args == "" {
		mudconn.SendFMessage("喔～那是我聽錯了。\r\n請問您的名字是：")
		cache.SetStep(0)
	} else if args == "y" {
		mudconn.SendFMessage(aurora.Sprintf("您使用%s做為您的名字。", aurora.Cyan(cache.User.ID)))
		cache.SetStep(2)
		mudconn.SendFMessage("請設定一個不容易被猜中的密碼：")
	} else {
		mudconn.SendFMessage(aurora.Sprintf("沒聽清楚您的回答。請問要使用%s做為你的名字嗎[%s/%s]？", aurora.Cyan(cache.User.ID), aurora.Cyan("Y"), aurora.Cyan("N")))
	}
	return
}

type GetPasswordCommand struct{} //2
func (c *GetPasswordCommand) Execute(pwd string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	cache.User.Password = pwd
	cache.SetStep(3)
	mudconn.SendFMessage("請再輸入一次密碼：")
	return
}

type GetConfirmPasswordCommand struct{} //3
func (c *GetConfirmPasswordCommand) Execute(cconfpwd string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	if cconfpwd != cache.User.Password {
		mudconn.SendFMessage("兩次輸入的密碼不同。\r\n請重新設定一個不容易被猜中的密碼：")
		cache.User.Password = ""
		cache.SetStep(2)
	} else {
		mudconn.SendFMessage("請取一個您在這個世界的化名，禁止使用不雅的化名，我們有權可以不經警告直接刪除使用不雅文字的角色。\r\n請問你的化名是？")
		cache.SetStep(4)
	}
	return
}

type GetCreatureNameCommand struct{} //4
func (c *GetCreatureNameCommand) Execute(name string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	if err = utility.ValidateStringLengRange(name, 1, 20); err != nil {
		return
	}
	cache.User.Name = name
	cache.SetStep(5)
	mudconn.SendFMessage("請選擇你的種族：")
	mudconn.SendFMessage(utility.ToCyan("Human"))
	return
}

type GetCreatureRaceCommand struct{} //5
func (c *GetCreatureRaceCommand) Execute(race string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	raceInt, err := Creature.ParseCreatureRace(race)
	if err != nil {
		mudconn.SendFMessage(err.Error())
		mudconn.SendFMessage("請選擇你的種族：")
		mudconn.SendFMessage(utility.ToCyan("Human"))
		err = nil
		return
	}
	cache.User.Race = raceInt
	cache.SetStep(6)
	mudconn.SendFMessage(aurora.Sprintf("您選擇%s作為您的種族", aurora.Cyan(raceInt.String())))
	mudconn.SendFMessage("您的性別：\r\n%s\r\n%s\r\n%s：", utility.ToCyan("Male"), utility.ToCyan("Female"), utility.ToCyan("Nature"))
	return
}

type GetSexCommand struct{} //6
func (c *GetSexCommand) Execute(args string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	args = strings.ToLower(args)
	switch args {
	case "m", "male":
		args = "Male"
	case "f", "female":
		args = "Female"
	case "n", "nature":
		args = "Nature"
	default:
		args = ""
	}
	if args == "" {
		mudconn.SendFMessage("請設定您的性別\r\n%s\r\n%s\r\n%s：", utility.ToCyan("Male"), utility.ToCyan("Female"), utility.ToCyan("Nature"))
	} else {
		cache.Step = 7
		cache.User.Gender = Player.ParseUserGender(args)
		cache.SetStep(7)
		mudconn.SendFMessage(aurora.Sprintf("您將性別設定為： %s", aurora.Cyan(cache.User.Gender.String())))
		mudconn.SendFMessage("請選擇您的職業：")
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Magic.String(), utility.ToCyan(Creature.Magic.StringEng())))
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Musician.String(), utility.ToCyan(Creature.Musician.StringEng())))
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Thief.String(), utility.ToCyan(Creature.Thief.StringEng())))
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Fighter.String(), utility.ToCyan(Creature.Fighter.StringEng())))
	}
	return
}

type GetCreatureCareerCommand struct{} //7
func (c *GetCreatureCareerCommand) Execute(career string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	career = strings.ToLower(career)
	careerInt := Creature.ParseCreatureCareer(career)
	if careerInt == -1 {
		mudconn.SendFMessage("沒聽清楚您的選擇，請再選擇一次您的職業：")
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Magic.String(), utility.ToCyan(Creature.Magic.StringEng())))
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Musician.String(), utility.ToCyan(Creature.Musician.StringEng())))
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Thief.String(), utility.ToCyan(Creature.Thief.StringEng())))
		mudconn.SendFMessage(aurora.Sprintf("%s(%s)", Creature.Fighter.String(), utility.ToCyan(Creature.Fighter.StringEng())))
	} else {
		mudconn.SendFMessage(aurora.Sprintf("您的職業為： %s", aurora.Cyan(cache.User.Career.String())))
		attr := careerInt.GetRandAttributes()
		mudconn.SendFMessage("您的屬性：")
		mudconn.SendFMessage(fmt.Sprintf("力量 - %d", attr.Strength))
		mudconn.SendFMessage(fmt.Sprintf("速度 - %d", attr.Dexterity))
		mudconn.SendFMessage(fmt.Sprintf("體格 - %d", attr.Constitution))
		mudconn.SendFMessage(fmt.Sprintf("智力 - %d", attr.Intelligent))
		mudconn.SendFMessage(fmt.Sprintf("知識 - %d", attr.Wisdom))
		mudconn.SendFMessage(aurora.Sprintf("這樣子的屬性你滿意嗎？[%s/%s]", aurora.Cyan("Y"), aurora.Cyan("N")))
		cache.User.AttributeBasic = attr
		cache.SetStep(99)
	}
	return
}

var m = new(sync.Mutex)

type ConfirmAttributeCommand struct{} //99
func (c *ConfirmAttributeCommand) Execute(ans string, mudconn *StructCollection.MudClient, cache *CacheService.LoginUser) (exitLoop bool, isCreated bool, err error) {
	ans = strings.ToLower(ans)
	if ans == "y" {
		cache.User.InitLv0()
		//存檔
		m.Lock()
		defer m.Unlock()
		tmpUser, _ := PlayerManagement.GetByID(cache.User.ID)
		if tmpUser == nil {
			err = PlayerManagement.Save(cache.User)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			exitLoop = true
			isCreated = true
			mudconn.SendFMessage("帳號%s建立完成。", aurora.Cyan(cache.User.ID))
		} else {
			mudconn.SendFMessage("您剛剛所選的帳號，被其它人先行一步用掉了，請重新建立帳號。")
			cache.SetStep(0)
		}
	} else {
		attr := cache.User.Career.GetRandAttributes()
		mudconn.SendFMessage("您的屬性：")
		mudconn.SendFMessage(fmt.Sprintf("力量 - %d", attr.Strength))
		mudconn.SendFMessage(fmt.Sprintf("速度 - %d", attr.Dexterity))
		mudconn.SendFMessage(fmt.Sprintf("體格 - %d", attr.Constitution))
		mudconn.SendFMessage(fmt.Sprintf("智力 - %d", attr.Intelligent))
		mudconn.SendFMessage(fmt.Sprintf("知識 - %d", attr.Wisdom))
		mudconn.SendFMessage(aurora.Sprintf("這樣子的屬性你滿意嗎？[%s/%s]", aurora.Cyan("Y"), aurora.Cyan("N")))
		cache.User.AttributeBasic = attr
		cache.SaveUser()
	}
	return
}

func init() {
	RegisterCommand(0, "", &IsNewCreationCommand{})
	RegisterCommand(1, "", &ConfirmCreationCommand{})
	RegisterCommand(2, "", &GetPasswordCommand{})
	RegisterCommand(3, "", &GetConfirmPasswordCommand{})
	RegisterCommand(4, "", &GetCreatureNameCommand{})
	RegisterCommand(5, "", &GetCreatureRaceCommand{})
	RegisterCommand(6, "", &GetSexCommand{})
	RegisterCommand(7, "", &GetCreatureCareerCommand{})
	RegisterCommand(99, "", &ConfirmAttributeCommand{})
}
