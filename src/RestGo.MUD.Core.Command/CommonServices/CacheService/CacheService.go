package CacheService

import (
	"github.com/patrickmn/go-cache"
)

type cacheStorage struct {
	inLoginProcessUser *cache.Cache
	allUsers           *cache.Cache
}

var pCacheInstance = new(cacheStorage)

// 新增一個Cache，也要加移除段落
func RemoveAllCacheByConnectionID(connectionID string) {
	CacheGetUser(connectionID).RemoveUser()
	p := userCache{}
	p.RemoveUser(connectionID)
}
