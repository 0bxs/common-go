package request

import (
	"bytes"
	"common/src/log"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"common/src/catch"

	"github.com/bytedance/sonic"
)

func PostJson[T, P any](url string, params P, fn func(header http.Header)) *T {
	buffer := bytes.NewBuffer(catch.Try1(sonic.Marshal(params)))
	request := catch.Try1(http.NewRequest(http.MethodPost, url, buffer))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
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

func PostFrom[T, P any](url0 string, params P, fn func(header http.Header)) *T {
	values := url.Values{}
	data := catch.Try1(sonic.Marshal(params))
	var m map[string]any
	catch.Try(sonic.Unmarshal(data, &m))
	for k, v := range m {
		values.Set(k, fmt.Sprint(v))
	}
	buffer := bytes.NewBuffer([]byte(values.Encode()))
	request := catch.Try1(http.NewRequest(http.MethodPost, url0, buffer))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

func PostFrom1[T, P any](url0 string, params map[string]P, fn func(header http.Header)) *T {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, fmt.Sprint(v))
	}
	buffer := bytes.NewBuffer([]byte(values.Encode()))
	request := catch.Try1(http.NewRequest(http.MethodPost, url0, buffer))
	log.Zap.Info(request)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
