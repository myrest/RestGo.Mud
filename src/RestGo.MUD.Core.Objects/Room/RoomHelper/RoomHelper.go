package RoomHelper

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
)

var Rooms = make(map[int]*Room.Room)

const DocumentRoomRoot = "Documents/Objects/Rooms"
