package Room

import (
	"encoding/json"
	"strings"
)

type CoverType int

// 沒有設定，預設行為模式為Natural
const (
	NoneCoverSetting CoverType = iota //未設定
	Natural                           //自然變化
	Dark                              //永夜
	Light                             //永晝
)

func (c CoverType) String() string {
	return [...]string{"未設定", "自然變化", "永夜", "永晝"}[c]
}

func (c *CoverType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*c = ParseCoverType(str)
	return nil
}

func ParseCoverType(arg string) CoverType {
	switch strings.ToLower(arg) {
	case "natural":
		return Natural
	case "dark":
		return Dark
	case "light":
		return Light
	default:
		return NoneCoverSetting
	}
}
