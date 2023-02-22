package ExitsPosition

import (
	"encoding/json"
	"strings"
)

type ExitName int

const (
	East ExitName = iota
	West
	South
	North
	Up
	Down
)

func (c *ExitName) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*c = ParseExitName(str)
	return nil
}

func (c ExitName) String() string {
	return [...]string{"東邊", "西邊", "南邊", "北邊", "上方", "下方"}[c]
}

func ParseExitName(arg string) ExitName {
	switch strings.ToLower(arg) {
	case "east":
		return East
	case "west":
		return West
	case "south":
		return South
	case "north":
		return North
	case "up":
		return Up
	default:
		return Down
	}
}
