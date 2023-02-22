package CacheService

//處理login中的User

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core/Creature/Player"
)

type LoginUser struct {
	User         Player.User
	Step         int
	LoginFailed  int
	connectionID string
}

var CacheGetUser = func(connectionID string) *LoginUser {
	if x, found := pCacheInstance.inLoginProcessUser.Get(connectionID); found {
		return x.(*LoginUser)
	}
	newcache := LoginUser{
		User:         Player.User{},
		Step:         0,
		connectionID: connectionID,
	}
	newcache.SaveUser()
	return &newcache
}

func (c *LoginUser) RemoveUser() {
	pCacheInstance.inLoginProcessUser.Delete(c.connectionID)
}

func (c *LoginUser) SetStep(step int) {
	c.Step = step
	c.SaveUser()
}

func (c *LoginUser) SetLoginFailed() {
	c.LoginFailed++
	c.SaveUser()
}

func (c *LoginUser) SaveUser() {
	pCacheInstance.inLoginProcessUser.Set(c.connectionID, c, cache.DefaultExpiration)
}

func init() {
	// Cache 10分鐘，11分鐘後清除
	pCacheInstance.inLoginProcessUser = cache.New(10*time.Minute, 10*time.Minute)
	var removeEvent = func(connectionID string, value interface{}) {
		msg := fmt.Sprintf("Login User Cache Connection ID: %s is removed.", connectionID)
		fmt.Println(msg)
	}
	pCacheInstance.inLoginProcessUser.OnEvicted(removeEvent)
}
