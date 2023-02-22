package Container

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
)

type ContainerPure struct {
	Items []BasicDefinition.IObjectBasic
}

// 實作IContainer
func (c *ContainerPure) GetOut(num int, name string) (obj BasicDefinition.IObjectBasic, err error) {
	//先確認物件是否存在
	index := -1
	for key, value := range c.Items {
		actionObject := value.GetObjectBasic()
		if strings.HasPrefix(strings.ToLower(actionObject.Name_EN), strings.ToLower(name)) {
			if num <= 1 {
				obj = value
				index = key
				break
			} else {
				num--
			}
		}
	}
	if index > -1 {
		if !c.Items[index].HaveCapability(BasicDefinition.CanBeMove) {
			err = status.Error(BasicDefinition.ObjectCannotMove, "object can't be moved")
		} else {
			c.remove(index)
		}
	} else {
		err = status.Error(codes.NotFound, "object not found")
	}
	return
}

func (c *ContainerPure) PutIn(obj BasicDefinition.IObjectBasic) {
	c.Items = append(c.Items, obj)
}

func (c *ContainerPure) remove(index int) {
	c.Items = append(c.Items[:index], c.Items[index+1:]...)
}

func (c *ContainerPure) GetObjPoniter(num int, name string) (obj BasicDefinition.IObjectBasic, err error) {
	//先確認物件是否存在
	index := -1
	for _, value := range c.Items {
		if strings.HasPrefix(strings.ToLower(value.GetObjectBasic().Name_EN), name) {
			if num <= 1 {
				obj = value
				return
			} else {
				num--
			}
		}
	}
	if index == -1 {
		err = fmt.Errorf("object not found")
	}
	return
}

type alignSymbol struct {
	Num              int
	Description_List string
	Name_EN          string
}

func (c *ContainerPure) ItemListForDisplay() string {
	var msgarr []string                          //放格式化好的陣列
	var maxNum = 1                               //重疊最大數字
	var maxLength int                            //物件名稱最長字數
	var mapItems = make(map[string]*alignSymbol) //統計資料用
	for _, item := range c.Items {
		obj := item.GetObjectBasic()
		mapKey := obj.Name_EN + obj.Name_CH
		if value, ok := mapItems[mapKey]; ok {
			value.Num++
			if value.Num > maxNum {
				maxNum = value.Num
			}
		} else {
			mapItems[mapKey] = &alignSymbol{
				Num:              1,
				Description_List: obj.Description_List,
				Name_EN:          obj.Name_EN,
			}
		}
		objNameLength := utf8.RuneCountInString(obj.Description_List)
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
		lineItem += item.Description_List + strings.Repeat(" ", maxLength-utf8.RuneCountInString(item.Description_List))
		lineItem += fmt.Sprintf(" (%s)", item.Name_EN)
		msgarr = append(msgarr, lineItem)
	}

	return strings.Join(msgarr, "\n")
}
