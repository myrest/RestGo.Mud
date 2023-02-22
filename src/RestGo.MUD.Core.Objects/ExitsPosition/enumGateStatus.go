package ExitsPosition

import "strings"

type DoorStatus int

const (
	Open DoorStatus = iota
	Close
	Locked
)

func (c DoorStatus) String() string {
	return [...]string{"開", "關", "鎖住"}[c]
}

func ParseGateStatus(arg string) DoorStatus {
	switch strings.ToLower(arg) {
	case "open":
		return Open
	case "close":
		return Close
	default:
		return Locked
	}
}
