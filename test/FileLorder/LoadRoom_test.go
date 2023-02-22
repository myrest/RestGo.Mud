package TestLoadObject

import (
	"fmt"
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room/RoomHelper"
)

func TestRoomLoader(t *testing.T) {
	const DocumentRoomRoot = "Documents/Rooms"
	if err := RoomHelper.LoadRoomsFromFolder(DocumentRoomRoot); err != nil {
		fmt.Println(err.Error())
	}

	if len(RoomHelper.Rooms) < 1 {
		t.Errorf("Load Object Got nothing.")
	}
}
