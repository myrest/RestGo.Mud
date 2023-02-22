package Room

import (
	"fmt"

	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

const lookDescription_layout = utility.RestString("{RoomTitle}\n明顯的出口有：{Exits}\n{Description}")

func (r *Room) LookDescription(flag ...bool) string {
	isAdmin := false
	if len(flag) > 0 && flag[0] {
		isAdmin = true
	}

	roomTitle := r.Title
	exits := ""
	for key := range r.Exits {
		exits += key.String() + " "
	}

	if isAdmin {
		roomTitle = fmt.Sprintf("%s(%d) ", r.Title, r.Geography.RoomID)
		exits = ""
		for key, value := range r.Exits {
			exits += fmt.Sprintf("%s(%d) ", key, value.RoomID)
		}
	}

	msg := lookDescription_layout

	msg.Replace("{RoomTitle}", roomTitle).
		Replace("{Exits}", exits).
		Replace("{Description}", r.Description)
	return msg.String()
}
