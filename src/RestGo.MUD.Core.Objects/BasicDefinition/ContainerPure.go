package BasicDefinition

import (
	"fmt"
	"strings"
)

type ContainerPure struct {
	Items []IObjectBasic
}

// 實作IContainer
func (c *ContainerPure) GetOut(num int, name string) (obj IObjectBasic, err error) {
	//先確認物件是否存在
	index := -1
	for key, value := range c.Items {
		if strings.HasPrefix(strings.ToLower(value.GetObjectBasic().Name_EN), name) {
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
		c.remove(index)
	} else {
		err = fmt.Errorf("你身上沒有那個東西。")
	}
	return
}

func (c *ContainerPure) PutIn(obj IObjectBasic) {
	c.Items = append(c.Items, obj)
}

func (c *ContainerPure) remove(index int) {
	c.Items = append(c.Items[:index], c.Items[index+1:]...)
}

func (c *ContainerPure) IsExist(num int, name string) bool {
	for _, value := range c.Items {
		if strings.HasPrefix(strings.ToLower(value.GetObjectBasic().Name_EN), name) {
			if num < 1 {
				return true
			} else {
				num--
			}
		}
	}
	return false
}

func (c *ContainerPure) GetObjPoniter(num int, name string) (obj IObjectBasic, err error) {
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
