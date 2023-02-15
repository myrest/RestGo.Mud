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

const itemList_layout = "%s(%s)"
const itemListAdmin_layout = "%s(%s)%s"

func (r *Room) ItemList(flag ...bool) string {
	layout := itemList_layout
	var msgarr []string
	if len(flag) > 0 && flag[0] {
		layout = itemListAdmin_layout
	}
	for _, item := range r.Items {
		obj := item.GetObjectBasic()
		msgarr = append(msgarr, fmt.Sprintf(layout, obj.Name_CH, obj.Name_EN, obj.ID[:5]))
	}

	return utility.AlignMessageByBrackets(msgarr)
}
