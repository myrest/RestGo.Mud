package PlayerManagement

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rest.com.tw/tinymud/src/RestGo.DataStorage/firebase"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/ObjectHelper"
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

	itemObjectBasicArr := ObjectHelper.GetItemsFromMapObject(doc["Items"])
	pUser.Items = itemObjectBasicArr

	return

}

func Save(pUser Player.User) (err error) {
	err = firebase.UpdateOrCreate(collection, pUser.ID, pUser)
	if err != nil {
		return fmt.Errorf("error Save: ID: %s, %s", pUser.ID, err)
	}
	return nil
}
