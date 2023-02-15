package Room

import "strings"

type GateStatus int

const (
	Open GateStatus = iota
	Close
	Locked
)

func (c GateStatus) String() string {
	return [...]string{"開", "關", "鎖住"}[c]
}

func ParseGateStatus(arg string) GateStatus {
	switch strings.ToLower(arg) {
	case "open":
		return Open
	case "close":
		return Close
	default:
		return Locked
	}
}
