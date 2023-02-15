package Player

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

func (u *User) Prompt() string {
	p := u.Settings.Prompt
	if p == "" {
		p = "*生命%h% 精神%m% 體力%v%*生命%H% 精神%M% 體力%V%*"
	}
	replacements := map[string]string{
		"%h%": strconv.Itoa(u.AttributeCurrent.HP),
		"%v%": strconv.Itoa(u.AttributeCurrent.MV),
		"%m%": strconv.Itoa(u.AttributeCurrent.Mana),
		"%H%": fmt.Sprintf("%v %%", utility.GetPercentage(u.AttributeBasic.HP, u.AttributeCurrent.HP)),
		"%V%": fmt.Sprintf("%v %%", utility.GetPercentage(u.AttributeBasic.MV, u.AttributeCurrent.MV)),
		"%M%": fmt.Sprintf("%v %%", utility.GetPercentage(u.AttributeBasic.Mana, u.AttributeCurrent.Mana)),
	}

	for key, value := range replacements {
		p = strings.ReplaceAll(p, key, value)
	}

	return p
}

func (u *User) LevelUP() (increaseHP int, increaseMV int, increaseMana int) {
	u.Level++
	increaseHP = (u.AttributeCurrent.Constitution * 20) + rand.Intn(u.AttributeCurrent.Constitution/2) + u.AttributeCurrent.Constitution/2
	u.AttributeBasic.HP = u.AttributeBasic.HP + increaseHP

	increaseMV = (u.AttributeCurrent.Constitution * 10) + rand.Intn(u.AttributeCurrent.Constitution/3) + u.AttributeCurrent.Constitution/3
	u.AttributeBasic.MV = u.AttributeBasic.MV + increaseMV

	if u.Career.CareerCategory() == Creature.Musician || u.Career.CareerCategory() == Creature.Magician {
		increaseMana = (u.AttributeCurrent.Intelligent * 10) + rand.Intn(u.AttributeCurrent.Intelligent/3) + u.AttributeCurrent.Intelligent/3
		u.AttributeBasic.Mana = u.AttributeBasic.Mana + increaseMana
	}
	return
}

func (u *User) InitLv0() {
	//todo:要補足其它屬性值(攻防)
	u.AttributeBasic.HP = 300
	u.AttributeBasic.MV = 300
	if u.Career.CareerCategory() == Creature.Musician {
		u.AttributeBasic.Mana = 300
	}
	u.AttributeCurrent = u.AttributeBasic.Copy()
	u.Settings.Prompt = "*生命%h% 精神%m% 體力%v%*"
}
