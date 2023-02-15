package Room

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ExitsPosition"
)

type Room struct {
	BasicDefinition.ContainerPure
	Title       string
	Ceiling     CoverType
	Description string
	Geography   StructCollection.WorldGeography
	Exits       map[ExitsPosition.ExitName]ExitsPosition.Exit
}
