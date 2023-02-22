package RoomHelper

import (
	"fmt"

	R "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
)

func GetRoom(RoomID int) (*R.Room, error) {
	r, ok := Rooms[RoomID]
	if ok {
		return r, nil
	} else {
		return nil, fmt.Errorf("房間代碼[%d]不存在", RoomID)
	}
}
