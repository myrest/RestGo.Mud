package PlayerManagement

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rest.com.tw/tinymud/src/RestGo.DataStorage/firebase"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
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

	itemObjectBasicArr := GetItemsFromMapObject(doc)
	pUser.Items = itemObjectBasicArr

	return

}

func GetItemsFromMapObject(doc map[string]interface{}) []BasicDefinition.IObjectBasic {
	var itemObjectBasicArr []BasicDefinition.IObjectBasic
	if value, ok := doc["Items"]; ok {
		if arr, ok := value.([]interface{}); ok {
			for i := 0; i < len(arr); i++ {
				mapObj, _ := utility.ConvertStructToMap(arr[i])
				//判斷要序列化成哪一種物件
				if _, ok := doc["Items"]; ok {
					//含有Items的key，為ContainerBasic
					var itemContainer BasicDefinition.ContainerBasic
					utility.ConvertMapToObject(mapObj, &itemContainer)
					innerItems := GetItemsFromMapObject(mapObj)
					if innerItems != nil {
						itemContainer.Items = innerItems
					}
					itemObjectBasicArr = append(itemObjectBasicArr, &itemContainer)
				} else {
					//ObjectBasic
					var itemBasic BasicDefinition.ObjectBasic
					utility.ConvertMapToObject(mapObj, &itemBasic)
					itemObjectBasicArr = append(itemObjectBasicArr, &itemBasic)
				}
			}
		}
	}
	return itemObjectBasicArr
}

func Save(pUser Player.User) (err error) {
	err = firebase.UpdateOrCreate(collection, pUser.ID, pUser)
	if err != nil {
		return fmt.Errorf("error Save: ID: %s, %s", pUser.ID, err)
	}
	return nil
}
