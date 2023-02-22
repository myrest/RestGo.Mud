package ObjectHelper

import (
	"fmt"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

func GetItemsFromMapObject(obj interface{}) []BasicDefinition.IObjectBasic {
	var rtn []BasicDefinition.IObjectBasic
	if obj == nil {
		return rtn
	}
	if arr, ok := obj.([]interface{}); ok {
		for i := 0; i < len(arr); i++ {
			mapObj, _ := utility.ConvertStructToMap(arr[i]) //將物件轉成map型式，才能以string的方式取出field
			rtn = append(rtn, GetIndividualItem(mapObj))
		}
	} else {
		mapObj, _ := utility.ConvertStructToMap(obj) //將物件轉成map型式，才能以string的方式取出field
		rtn = append(rtn, GetIndividualItem(mapObj))
	}
	return rtn
}

func GetIndividualItem(mapObj map[string]interface{}) (rtn BasicDefinition.IObjectBasic) {
	objType := BasicDefinition.ParseObjectType(fmt.Sprint(mapObj["ObjectType"])) //取出物件類別
	switch objType {
	case BasicDefinition.Container:
		item := Container.ContainerObject{}
		utility.ConvertMapToObject(mapObj, &item)
		innerItems := GetItemsFromMapObject(mapObj["Items"])
		item.Items = innerItems
		return
	default:
		var item BasicDefinition.ObjectBasic
		utility.ConvertMapToObject(mapObj, &item)
		return
	}
}
