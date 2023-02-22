package RoomHelper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ExitsPosition"
	R "rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/Room"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Config"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

// 從檔案讀案，放到Cache Start//////////////////////////////////////////////////////////////
type roomJson struct {
	R.Room
	RoomID int
	Exits  map[string]ExitsPosition.Exit
	Extra  json.RawMessage
}

type roomsJson struct {
	Rooms          []roomJson
	DefaultCeiling R.CoverType
}

func (j *roomJson) ToRoom(World string, Region string, Area string, DefaultCeiling R.CoverType) (rtn *R.Room) {
	if DefaultCeiling == R.NoneCoverSetting {
		DefaultCeiling = R.Natural
	}
	if j.Ceiling == R.NoneCoverSetting {
		j.Ceiling = DefaultCeiling
	}
	rtn = &j.Room
	rtn.Geography.World = World
	rtn.Geography.Region = Region
	rtn.Geography.Area = Area
	rtn.Geography.RoomID = j.RoomID
	rtn.Description = utility.InsertNewLine(rtn.Description, Config.ServiceConfig.MaxLength)
	rtn.Exits = make(map[ExitsPosition.ExitName]ExitsPosition.Exit)
	for k, j := range j.Exits { //Todo:出口要排序，顯示出來心比較好看
		rtn.Exits[ExitsPosition.ParseExitName(k)] = j
	}

	// 处理额外字段
	var extra map[string]interface{}
	err := json.Unmarshal(j.Extra, &extra)
	if err != nil {
		fmt.Println("RoomID", rtn.Geography.RoomID, "有未處理到的屬性，欄位未知。")
		return
	} else {
		fmt.Println("RoomID", rtn.Geography.RoomID, "有未處理到的屬性。")
		unHandelAttr := ""
		for k := range extra {
			unHandelAttr += k + ","
		}
	}

	return rtn
}

func (rs *roomsJson) ToRoomCache(World string, Region string, Area string) error {
	for _, room := range rs.Rooms {
		if err := AddRoom(room.ToRoom(World, Region, Area, rs.DefaultCeiling)); err != nil {
			return err
		}
	}
	return nil
}

var documentRoomRoot string

func LoadRoomsFromFolder(DocumentRoomRoot string) error {
	documentRoomRoot = DocumentRoomRoot
	err := filepath.Walk(DocumentRoomRoot, loadRoomsFromFileWithCheck)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", DocumentRoomRoot, err)
	}
	return err
}

func loadRoomsFromFileWithCheck(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && strings.HasSuffix(path, ".json") {
		pathInformation := strings.Replace(path, documentRoomRoot, "", 1)
		// 將路徑分割成三個部分，分別代表 World、Region、Area
		parts := strings.Split(pathInformation, string(os.PathSeparator))
		//只處理第三層檔案
		if len(parts) != 5 {
			return nil
		}
		world := parts[1]
		region := parts[2]
		area := parts[3]

		var rooms = new(roomsJson)
		if err := utility.UnmarshalJsonFile(path, rooms); err != nil {
			panic(err)
		}
		if err := rooms.ToRoomCache(world, region, area); err != nil {
			return err
		}
	}
	return nil
}

func AddRoom(room *R.Room) error {
	if _, ok := Rooms[room.Geography.RoomID]; ok {
		return fmt.Errorf("房間重覆定義[%d]", room.Geography.RoomID)
	} else {
		Rooms[room.Geography.RoomID] = room
		return nil
	}
}

// 從檔案讀案，放到Cache 結束//////////////////////////////////////////////////////////////
