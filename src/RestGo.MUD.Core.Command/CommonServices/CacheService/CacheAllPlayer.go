package CacheService

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/ScheduleService"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Management/PlayerManagement"
)

// 處理登入成功的User Cache
var WorldUsersList = make(map[Creature.Geography][]*StructCollection.MudClient)

// 每10分鐘Save一次Cache，同時要Save到firebase上
func reFlashOnLineLoginedUser() {
	msg := fmt.Sprintf("Job A: %v", time.Now())
	fmt.Println(msg)
	//Checking: 假設Key值的Geography的RoomeID，與User.RoomID是相同的
	for _, MudClients := range WorldUsersList {
		if len(MudClients) > 0 {
			for _, client := range MudClients {
				connectionID := client.ConnectionID
				err := PlayerManagement.Save(client.User)
				if err != nil {
					fmt.Println("使用者存檔失敗。" + err.Error())
				}
				p := userCache{}
				p.SaveUser(connectionID) //避免20分鐘沒動作就被從Cache刪掉
			}
		}
	}
}

type userCache struct {
	Player.User
}

var CacheGetFromALLUser = func(connectionID string) *userCache {
	if x, found := pCacheInstance.allUsers.Get(connectionID); found {
		return x.(*userCache)
	}
	user := userCache{}
	user.SaveUser(connectionID)
	return &user
}

func (c *userCache) SaveUser(connectionID string) {
	pCacheInstance.allUsers.Set(connectionID, c, cache.DefaultExpiration)
}

func (c *userCache) RemoveUser(connectionID string) {
	pCacheInstance.allUsers.Delete(connectionID)
}

func init() {
	// Cache 10分鐘，20分鐘後清除
	pCacheInstance.allUsers = cache.New(10*time.Minute, 20*time.Minute)
	var removeEvent = func(connectionID string, value interface{}) {
		msg := fmt.Sprintf("ALL User Cache Connection ID: %s is removed.", connectionID)
		fmt.Println(msg)
	}
	pCacheInstance.allUsers.OnEvicted(removeEvent)

	//每10分鐘Save一次AllUser Cache
	ScheduleService.NewTimerInMins("定時存檔", 10, reFlashOnLineLoginedUser)
}
