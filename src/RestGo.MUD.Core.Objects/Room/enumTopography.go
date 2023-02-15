package Room

import "strings"

type Topography int

const (
	Normal   Topography = iota //平地
	Beach                      //海水、溪流
	Mountain                   //山地
)

func (c Topography) String() string {
	return [...]string{"平地", "海水、溪流", "山地"}[c]
}

func ParseTopography(arg string) Topography {
	switch strings.ToLower(arg) {
	case "normal":
		return Normal
	case "beach":
		return Beach
	default:
		return Mountain
	}
}
