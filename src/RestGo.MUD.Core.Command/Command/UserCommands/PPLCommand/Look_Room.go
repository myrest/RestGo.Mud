package PPLCommand

import (
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
)

func (c *LookCommand) LookRoom(room *Room.Room) (msg string) {
	//房間描述
	msg = room.LookDescription()
	itemList := room.ContainerPure.ItemListForDisplay()
	if itemList != "" {
		msg = msg + "\n" + itemList
	}
	return
}
