package white_list

import (
	"fmt"
	"time"

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
		return redis.GetI64(key(id))
	})
}

func Set(id, expireTime int64) {
	redis.SetPx(key(id), expireTime, expire)
	userCache.Set(id, expireTime, expire)
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
