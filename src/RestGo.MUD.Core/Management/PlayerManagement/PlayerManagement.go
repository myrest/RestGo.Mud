package PlayerManagement

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rest.com.tw/tinymud/src/RestGo.DataStorage/firebase"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/Container"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

var collection = "User"

func GetByID(id string) (pUser *Player.User, err error) {
	doc, err := firebase.GetByID(collection, id)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			fmt.Println("User not found: ", id)
		} else {
			fmt.Println("status.Code(err): ", status.Code(err))
			fmt.Println("message: ", err.Error())
		}
		return
	}

	utility.ConvertMapToObject(doc, &pUser)

	itemObjectBasicArr := GetItemsFromMapObject(doc["Items"])
	pUser.Items = itemObjectBasicArr

	return

}

func GetItemsFromMapObject(obj interface{}) []BasicDefinition.IObjectBasic {
	var rtn []BasicDefinition.IObjectBasic
	if arr, ok := obj.([]interface{}); ok {
		for i := 0; i < len(arr); i++ {
			mapObj, _ := utility.ConvertStructToMap(arr[i])                              //將物件轉成map型式，才能以string的方式取出field
			objType := BasicDefinition.ParseObjectType(fmt.Sprint(mapObj["ObjectType"])) //取出物件類別
			switch objType {
			case BasicDefinition.Container:
				item := Container.ContainerObject{}
				utility.ConvertMapToObject(mapObj, &item)
				innerItems := GetItemsFromMapObject(mapObj["Items"])
				item.Items = innerItems
				rtn = append(rtn, &item)
			default:
				var item BasicDefinition.ObjectBasic
				utility.ConvertMapToObject(mapObj, &item)
				rtn = append(rtn, &item)
			}
		}
	}
	return rtn
}

func Save(pUser Player.User) (err error) {
	err = firebase.UpdateOrCreate(collection, pUser.ID, pUser)
	if err != nil {
		return fmt.Errorf("error Save: ID: %s, %s", pUser.ID, err)
	}
	return nil
}
