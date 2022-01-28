package util

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var cacheInstanceMap map[string]cache.Cache = make(map[string]cache.Cache)

func Set(bigKey string, smallKey string, value string, expire int) {

	var cacheInstance cache.Cache
	cacheInstance, ok := cacheInstanceMap[bigKey]

	if !ok {
		cacheInstance = *cache.New(5*time.Minute, 10*time.Minute)
		cacheInstanceMap[bigKey] = cacheInstance
	}

	cacheInstance.Set(smallKey, value, time.Duration(expire*1000*1000*1000))
}

func Get(bigKey string, smallKey string) (value string, result bool) {
	var cacheInstance cache.Cache
	cacheInstance, ok := cacheInstanceMap[bigKey]
	if ok {
		if x, found := cacheInstance.Get(smallKey); found {
			return x.(string), true
		}
	}
	return "", false
}
