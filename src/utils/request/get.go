package request

import (
	"io"
	"net/http"
	"strings"

	"github.com/0bxs/common-go/src/catch"
	"github.com/0bxs/common-go/src/collection/dict"

	"github.com/bytedance/sonic"
)

func Get1[T any](url0 string, params dict.Dict[string, string]) *T {
	builder := strings.Builder{}
	builder.WriteString(url0)
	if params.Len() > 0 {
		builder.WriteByte('?')
		flag := true
		params.ForEach(func(fieldName string, value string) {
			if flag {
				flag = false
			} else {
				builder.WriteByte('&')
			}
			builder.WriteString(fieldName)
			builder.WriteByte('=')
			builder.WriteString(value)
		})
	}
	body := catch.Try1(http.Get(builder.String())).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	t := new(T)
	bodyBytes := catch.Try1(io.ReadAll(body))
	catch.Try(sonic.Unmarshal(bodyBytes, t))
	return t
}

func Get0[T any](url string, fn func(header http.Header)) *T {
	request := catch.Try1(http.NewRequest(http.MethodGet, url, nil))
	fn(request.Header)
	body := catch.Try1(http.DefaultClient.Do(request)).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	t := new(T)
	bodyBytes := catch.Try1(io.ReadAll(body))
	catch.Try(sonic.Unmarshal(bodyBytes, t))
	return t
}
