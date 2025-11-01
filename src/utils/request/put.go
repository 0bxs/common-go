package request

import (
	"io"
	"net/http"

	"github.com/0bxs/common-go/src/catch"

	"github.com/bytedance/sonic"
)

func Put[T any](url string, body io.Reader, fn func(h http.Header)) *T {
	req := catch.Try1(http.NewRequest(http.MethodPut, url, body))
	fn(req.Header)
	respBody := catch.Try1(http.DefaultClient.Do(req)).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(respBody)
	t := new(T)
	data := catch.Try1(io.ReadAll(respBody))
	if len(data) == 0 {
		return nil
	}
	catch.Try(sonic.Unmarshal(data, t))
	return t
}
