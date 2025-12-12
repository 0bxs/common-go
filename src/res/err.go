package res

import (
	"fmt"
	"net/http"

	"github.com/0bxs/common-go/src/log"
	"github.com/0bxs/common-go/src/status"
	"github.com/0bxs/common-go/src/utils/rawJson"
	"github.com/gofiber/fiber/v2"
)

var systemErrorJson = rawJson.RawJson(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, status.SystemError, "系统繁忙，请稍后再试"))

type Err struct {
	Code int8   `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (r Err) Error() string {
	if r.Msg == "" {
		return fmt.Sprintf(`{"code":%d}`, r.Code)
	}
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, r.Code, r.Msg)
}

func ErrorHandler0(ctx *fiber.Ctx) error {
	defer func() {
		switch e := recover().(type) {
		case nil:
		case Err:
			log.Zap.Error(e)
			_ = ctx.JSON(e)
		default:
			log.Zap.Error(e)
			ctx.Status(http.StatusInternalServerError)
			_ = ctx.JSON(systemErrorJson)
		}
	}()
	_ = ctx.Next()
	return nil
}
