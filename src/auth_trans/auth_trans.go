package auth_trans

import (
	"fmt"
	"time"

	"github.com/0bxs/common-go/src/collection/bit_map"
	"github.com/0bxs/common-go/src/collection/dict"
	"github.com/0bxs/common-go/src/collection/set"
	"github.com/0bxs/common-go/src/collection/vec"
	"github.com/0bxs/common-go/src/redis"
	"github.com/0bxs/common-go/src/utils/cache"
	"github.com/0bxs/common-go/src/utils/option"
)

var (
	redisKey  = ""
	userCache = cache.New[int64, set.Set[int16]](1000)
	expire    = int64(time.Second * 60 * 60 * 24 * 7)
)

func Get(id int64) option.Opt[set.Set[int16]] {
	return userCache.Get(id).Else(func() option.Opt[set.Set[int16]] {
		opt := redis.GetString(key(id))
		if opt.IsSome() {
			bitMap := bit_map.BytesBitMapNew([]byte(opt.V))
			v := bitMap.ToSet()
			userCache.Set(id, v, expire)
			return option.Some(v)
		}
		return option.None[set.Set[int16]]()
	})
}

func Gets(ids vec.Vec[int64]) (dict.Dict[int64, set.Set[int16]], vec.Vec[int64]) {
	missIds := vec.New[int64](ids.Len())
	d := dict.New[int64, set.Set[int16]](ids.Len())
	ids.ForEach(func(userId int64) {
		userCache.Get(userId).MapOrElse(func() {
			missIds.Append(userId)
		}, func(s set.Set[int16]) {
			d.Store(userId, s)
		})
	})
	return d, missIds
}

func Set(id int64, auth bit_map.BytesBitMap) {
	authSet := auth.ToSet()
	authStr := auth.ToStr()
	Get(id).Map(func(t set.Set[int16]) {
		authSet = authSet.Or(t)
		var v []byte
		authSet.ForEach(func(u int16) {
			v = append(v, byte(u))
		})
		authStr = string(v)
	})
	redis.Set(key(id), authStr)
	userCache.Set(id, authSet, expire)
}

func GetDel(id int64) option.Opt[set.Set[int16]] {
	opt := redis.GetBytesDel(key(id))
	if opt.IsSome() {
		userCache.Del(id)
		bitMap := bit_map.BytesBitMapNew(opt.V)
		return option.Some(bitMap.ToSet())
	}
	return option.None[set.Set[int16]]()
}

func Del(id int64) {
	redis.Del(key(id))
	userCache.Del(id)
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
