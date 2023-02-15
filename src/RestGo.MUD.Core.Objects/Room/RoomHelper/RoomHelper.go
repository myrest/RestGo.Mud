package RoomHelper

import (
	"fmt"

	R "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
)

var Rooms = make(map[int]*R.Room)

const DocumentRoomRoot = "Documents/Objects/Rooms"

func init() {
	err := loadRoomsFromFolder()
	if err != nil {
		fmt.Println(err.Error())
	}
}
