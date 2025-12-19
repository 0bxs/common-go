package redis

import (
	"context"

	"github.com/0bxs/common-go/src/catch"
	"github.com/0bxs/common-go/src/collection/vec"
	"github.com/0bxs/common-go/src/types"

	"github.com/bytedance/sonic"
)

func LPush[T string | types.Number](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "LPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = value
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func LPushObj[T any](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "LPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = catch.Try1(sonic.Marshal(value))
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func RPush[T string | types.Number](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "RPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = value
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func RPushObj[T any](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "RPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = catch.Try1(sonic.Marshal(value))
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func LRange[T string | types.Number](key string, start, stop int64) vec.Vec[T] {
	result := catch.Try1(client.Do(context.Background(), "LRANGE", key, start, stop).Slice())
	temp := vec.New[T](len(result))
	for _, v := range result {
		temp.Append(v.(T))
	}
	return temp
}

func LRangeObj[T any](key string, start, stop int64) vec.Vec[T] {
	result := catch.Try1(client.Do(context.Background(), "LRANGE", key, start, stop).Slice())
	temp := vec.New[T](len(result))
	for _, v := range result {
		t := new(T)
		catch.Try(sonic.Unmarshal([]byte(v.(string)), t))
		temp.Append(*t)
	}
	return temp
}

func LRangeObj1[T any](key string, start, stop int64) (vec.Vec[T], vec.Vec[string]) {
	result := catch.Try1(client.Do(context.Background(), "LRANGE", key, start, stop).Slice())
	temp := vec.New[T](len(result))
	temp1 := vec.New[string](len(result))
	for _, v := range result {
		t := new(T)
		tempStr := v.(string)
		temp1.Append(tempStr)
		catch.Try(sonic.Unmarshal([]byte(tempStr), t))
		temp.Append(*t)
	}
	return temp, temp1
}

func LRem[T string | types.Number](key string, value T) {
	catch.Try(client.Do(context.Background(), "LREM", key, 1, value).Err())
}

func LRemObj[T any](key string, value T) {
	catch.Try(client.Do(context.Background(), "LREM", key, 1, string(catch.Try1(sonic.Marshal(value)))).Err())
}

func LLen(key string) int {
	return catch.Try1(client.Do(context.Background(), "LLEN", key).Int())
}
