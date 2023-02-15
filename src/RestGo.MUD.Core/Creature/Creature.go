package Creature

import (
	"encoding/json"
	"fmt"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
)

type Creature struct {
	ID               string
	Name             string
	Description      string
	Title            string
	AttributeBasic   UserAttributes
	AttributeCurrent UserAttributes
	Career           CreatureCareer
	Race             CreatureRace
	RoomID           int
	Level            int
	BasicDefinition.ContainerPure
}

type UserAttributes struct {
	HP           int
	MV           int
	Mana         int
	Constitution int
	Strength     int
	Dexterity    int
	Wisdom       int
	Intelligent  int
}

func (a *UserAttributes) Copy() UserAttributes {
	rtn := UserAttributes{}
	jsonStr, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(jsonStr, &rtn)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rtn
}

type Geography struct {
	RoomID   int
	AreaID   int
	RegionID int
	WordID   int
}
