package trans

import (
	"github.com/0bxs/common-go/src/catch"

	"github.com/bytedance/sonic"
)

func Obj2Map[T any](obj any) map[string]T {
	data := catch.Try1(sonic.Marshal(obj))
	temp := new(map[string]T)
	catch.Try(sonic.Unmarshal(data, temp))
	return *temp
}

func Map2Obj[T any](m map[string]any) T {
	data := catch.Try1(sonic.Marshal(m))
	t := new(T)
	catch.Try(sonic.Unmarshal(data, t))
	return *t
}
