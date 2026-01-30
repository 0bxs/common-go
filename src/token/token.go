package token

import (
	"fmt"
	"time"

	"github.com/0bxs/common-go/src/function"
	"github.com/0bxs/common-go/src/redis"
	"github.com/0bxs/common-go/src/utils/option"
	"github.com/Code-Hex/go-generics-cache"
	"github.com/Code-Hex/go-generics-cache/policy/lru"
)

var (
	redisKey  = ""
	userCache = new(cache.Cache[int64, int64])
	expireI64 = int64(time.Second * 60 * 60 * 24 * 7)
	expire    = cache.WithExpiration(time.Second * 60 * 60 * 24 * 7)
)

func Get(id int64) option.Opt[int64] {
	return option.OptOf(userCache.Get(id)).Else(func() option.Opt[int64] {
		opt := redis.GetI64(key(id))
		opt.Map(func(t int64) {
			userCache.Set(id, t, expire)
		})
		return opt
	})
}

func Set(id, expireTime int64) {
	redis.SetPx(key(id), expireTime, expireI64)
	userCache.Set(id, expireTime, expire)
}

func Del(id int64) {
	redis.Del(key(id))
	userCache.Delete(id)
}

func KickOut(ids []int64) {
	redis.Del(function.Map(ids, func(t int64) string {
		userCache.Delete(t)
		return key(t)
	})...)

}

func key(id int64) string {
	return fmt.Sprintf(redisKey, id)
}

func Init(key0 string, expire0 int64, cap int) {
	redisKey = key0
	if expire0 > 0 {
		expireI64 = expire0
		expire = cache.WithExpiration(time.Duration(expire0))
	}
	userCache = cache.New(cache.AsLRU[int64, int64](lru.WithCapacity(cap)))
}
