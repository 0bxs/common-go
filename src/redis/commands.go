package redis

import (
	"context"

	"common/src/catch"
)

func Exists(key string) bool {
	result := catch.Try1(client.Exists(context.Background(), key).Result())
	if result == 1 {
		return true
	}
	return false
}

func Del(keys ...string) {
	catch.Try(client.Del(context.Background(), keys...).Err())
}

func TTL(key string) int64 {
	result := catch.Try1(client.TTL(context.Background(), key).Result())
	return result.Milliseconds()
}

func PTtl(key string) int64 {
	result := catch.Try1(client.PTTL(context.Background(), key).Result())
	return result.Milliseconds()
}
