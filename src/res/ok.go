package res

import (
	"fmt"

	"github.com/0bxs/common-go/src/status"
	"github.com/0bxs/common-go/src/utils/rawJson"
	"github.com/gofiber/fiber/v2"
)

type Ok[T any] struct {
	Code int8 `json:"code"`
	Data T    `json:"data"`
}

type Res[T any] struct {
	Code int8   `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data"`
}

var (
	okJson = rawJson.RawJson(fmt.Sprintf(`{"code":%d}`, status.OK))
)

func Ok0(ctx *fiber.Ctx) error {
	return ctx.JSON(okJson)
}

func Ok1[T any](ctx *fiber.Ctx, data T) error {
	return ctx.JSON(Ok[T]{Code: status.OK, Data: data})
}
