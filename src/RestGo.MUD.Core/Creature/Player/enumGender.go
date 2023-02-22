package Player

import "strings"

type UserGender int

const (
	Male UserGender = iota
	Female
	Nature
)

func (c UserGender) String() string {
	return [...]string{"男性", "女性", "中性"}[c]
}

func ParseUserGender(gender string) UserGender {
	switch strings.ToLower(gender) {
	case "male":
		return Male
	case "female":
		return Female
	default:
		return Nature
	}
}
