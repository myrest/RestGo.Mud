package BasicDefinition

import "strings"

type BasicCapability int

const (
	CanBeMove BasicCapability = iota
	CanBeEat
	CanBeFill
)

type ObjectType int

const (
	Basic ObjectType = iota
	Container
	Fixed
)

func (c ObjectType) String() string {
	return [...]string{"Container", "Fixed"}[c]
}

func ParseObjectType(objType string) ObjectType {
	switch strings.ToLower(objType) {
	case "Container":
		return Container
	case "Fixed":
		return Fixed
	default:
		return Container
	}
}
