package Room

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

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

type alignSymbol struct {
	Num     int
	Name_CH string
	Name_EN string
}

func (r *Room) ItemList() string {
	var msgarr []string                          //放格式化好的陣列
	var maxNum = 1                               //重疊最大數字
	var maxLength int                            //物件名稱最長字數
	var mapItems = make(map[string]*alignSymbol) //統計資料用
	for _, item := range r.Items {
		obj := item.GetObjectBasic()
		mapKey := obj.Name_EN + obj.Name_CH
		if value, ok := mapItems[mapKey]; ok {
			value.Num++
			if value.Num > maxNum {
				maxNum = value.Num
			}
		} else {
			mapItems[mapKey] = &alignSymbol{
				Num:     1,
				Name_CH: obj.Name_CH,
				Name_EN: obj.Name_EN,
			}
		}
		objNameLength := utf8.RuneCountInString(obj.Name_CH)
		if objNameLength > maxLength {
			maxLength = objNameLength
		}
	}

	maxNumLength := len(strconv.Itoa(maxNum))
	for _, item := range mapItems {
		var lineItem string
		if maxNum > 1 {
			itemNumLength := len(strconv.Itoa(item.Num))
			//要加數量
			if item.Num > 1 {
				//在左邊加空白
				lineItem = fmt.Sprintf("%s(%d)", strings.Repeat(" ", maxNumLength-itemNumLength), item.Num)
			} else {
				lineItem = strings.Repeat(" ", maxNum+2)
			}
		}
		lineItem += item.Name_CH + strings.Repeat(" ", maxLength-utf8.RuneCountInString(item.Name_CH))
		lineItem += fmt.Sprintf(" (%s)", item.Name_EN)
		msgarr = append(msgarr, lineItem)
	}

	return strings.Join(msgarr, "\n")
}
