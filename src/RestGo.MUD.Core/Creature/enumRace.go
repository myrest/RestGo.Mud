package Creature

import (
	"fmt"
)

type CreatureRace int

const (
	Human = 0
)

func (c CreatureRace) String() string {
	return [...]string{"人類"}[c]
}

func ParseCreatureRace(s string) (CreatureRace, error) {
	switch s {
	case "human", "h":
		return Human, nil
	default:
		return Human, fmt.Errorf("沒有找到名為 %s 的種族", s)
	}
}
