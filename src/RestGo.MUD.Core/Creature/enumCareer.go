package Creature

import (
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type CreatureCareer int

const (
	Thief CreatureCareer = iota
	Ninja
	Assassin
	Ranger
	Fighter
	Warrior
	Knight
	Paladin
	Sage
	Wizard
	Magic
	Magician
	Archmage
	Musician
	Bard
	Chief
)

func (c CreatureCareer) String() string {
	return [...]string{"盜賊", "忍者", "殺手", "流浪人", "戰士", "武士", "騎士", "俠客", "贀者", "巫師", "施法者", "魔法使", "大法師", "音樂家", "吟遊詩人", "領袖"}[c]
}

func (c CreatureCareer) StringEng() string {
	return [...]string{"Thief", "Ninja", "Assassin", "Ranger", "Fighter", "Warrior", "Knight", "Paladin", "Sage", "Wizard", "Magic", "Magician", "Archmage", "Musician", "Bard", "Chief"}[c]
}

func (c CreatureCareer) CareerCategory() CreatureCareer {
	return [...]CreatureCareer{Thief, Thief, Thief, Thief, Fighter, Fighter, Fighter, Fighter, Magician, Magician, Magician, Magician, Musician, Musician, Musician, Musician}[c]
}

func ParseCreatureCareer(input string) CreatureCareer {
	switch strings.ToLower(input) {
	case "t", "thief":
		return Thief
	case "n", "ninja":
		return Ninja
	case "a", "assassin":
		return Assassin
	case "r", "ranger":
		return Ranger
	case "f", "fighter":
		return Fighter
	case "w", "warrior":
		return Warrior
	case "k", "knight":
		return Knight
	case "p", "paladin":
		return Paladin
	case "s", "sage":
		return Sage
	case "z", "wizard":
		return Wizard
	case "m", "magic":
		return Magic
	case "mg", "magician":
		return Magician
	case "am", "archmage":
		return Archmage
	case "mu", "musician":
		return Musician
	case "b", "bard":
		return Bard
	case "c", "chief":
		return Chief
	default:
		return -1
	}
}

func (a *CreatureCareer) getMaxAttribute() map[string]int {
	switch a.StringEng() {
	case "Thief":
		return map[string]int{"Strength": 18, "Intelligent": 16, "Wisdom": 18, "Dexterity": 20, "Constitution": 18}
	case "Magic":
		return map[string]int{"Strength": 14, "Intelligent": 20, "Wisdom": 20, "Dexterity": 18, "Constitution": 14}
	case "Musician":
		return map[string]int{"Strength": 16, "Intelligent": 18, "Wisdom": 20, "Dexterity": 18, "Constitution": 16}
	case "Fighter":
		return map[string]int{"Strength": 20, "Intelligent": 14, "Wisdom": 14, "Dexterity": 14, "Constitution": 20}
	default:
		return map[string]int{"Strength": 8, "Intelligent": 8, "Wisdom": 8, "Dexterity": 8, "Constitution": 8}
	}
}

func (a *CreatureCareer) GetRandAttributes() UserAttributes {
	att := a.getRandAttributesWithExtraPoints()
	return UserAttributes{
		Constitution: att["Constitution"],
		Strength:     att["Strength"],
		Dexterity:    att["Dexterity"],
		Wisdom:       att["Wisdom"],
		Intelligent:  att["Intelligent"],
	}
}

func (a *CreatureCareer) getRandAttributesWithExtraPoints() map[string]int {
	maxAtt := a.getMaxAttribute()

	// 隨機取得extraPoint
	extraPoint := rand.Intn(8) + 25

	// 設置attributes map
	attributes := make(map[string]int)

	//初始化屬性點
	for attrName := range maxAtt {
		attributes[attrName] = 8
	}

	for extraPoint > 0 {
		attName := a.getRandAttrString()
		if attributes[attName] < maxAtt[attName] {
			attributes[attName]++
			extraPoint--
		}
		//防呆，防止所有屬性都滿點，會變成無窮迴圈
		if reflect.DeepEqual(maxAtt, attributes) {
			break
		}
	}
	return attributes
}

func (a *CreatureCareer) getRandAttrString() string {
	i := rand.Intn(5)
	return attributeMap[i]
}

var attributeMap = map[int]string{
	1: "Constitution",
	2: "Strength",
	3: "Dexterity",
	4: "Wisdom",
	0: "Intelligent",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
