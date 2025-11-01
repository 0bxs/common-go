package blacklist

import (
	"fmt"

	"common/src/redis"
	"common/src/utils/cache"
)

var (
	redisKey  = ""
	userCache = cache.New[int64, bool](1000)
)

func Get(id int64, expireTime int64) bool {
	return userCache.Get(id + expireTime).GetElse(func() bool {
		key0 := key(id, expireTime)
		exists := redis.Exists(key0)
		if exists {
			userCache.Set(id+expireTime, true, redis.PTtl(key0))
		}
		return exists
	})
}

func Set(id int64, currentTime, oldExpireTime int64) {
	if oldExpireTime > currentTime {
		expire := oldExpireTime - currentTime
		redis.SetPx(key(id, oldExpireTime), "", expire)
		userCache.Set(id+oldExpireTime, true, expire)
	}
}

func key(id int64, expireTime int64) string {
	return fmt.Sprintf(redisKey, id, expireTime)
}

func Init(key0 string) {
	redisKey = key0
}
