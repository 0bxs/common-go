package token

import (
	"fmt"
	"time"

	"github.com/0bxs/common-go/src/function"
	"github.com/0bxs/common-go/src/redis"
	"github.com/0bxs/common-go/src/utils/cache"
	"github.com/0bxs/common-go/src/utils/option"
)

var (
	redisKey  = ""
	userCache = cache.New[int64, int64](1000)
	expire    = int64(time.Second * 60 * 60 * 24 * 7)
)

func Get(id int64) option.Opt[int64] {
	return userCache.Get(id).Else(func() option.Opt[int64] {
		opt := redis.GetI64(key(id))
		opt.Map(func(t int64) {
			userCache.Set(id, t, expire)
		})
		return opt
	})
}

func Set(id, expireTime int64) {
	redis.SetPx(key(id), expireTime, expire)
	userCache.Set(id, expireTime, expire)
}

func Del(id int64) {
	redis.Del(key(id))
	userCache.Del(id)
}

func KickOut(ids []int64) {
	redis.Del(function.Map(ids, func(t int64) string {
		userCache.Del(t)
		return key(t)
	})...)

}

func key(id int64) string {
	return fmt.Sprintf(redisKey, id)
}

func Init(key0 string, expire0 int64) {
	redisKey = key0
	if expire0 > 0 {
		expire = expire0
	}
}
