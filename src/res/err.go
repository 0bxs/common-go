package res

import (
	"fmt"
)

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
