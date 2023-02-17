package Room

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ExitsPosition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
)

type Room struct {
	Container.ContainerPure
	Title       string
	Ceiling     CoverType
	Description string
	Geography   StructCollection.WorldGeography
	Exits       map[ExitsPosition.ExitName]ExitsPosition.Exit
}
