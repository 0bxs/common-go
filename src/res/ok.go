package res

type Ok[T any] struct {
	Code int8 `json:"code"`
	Data T    `json:"data"`
}

type Res[T any] struct {
	Code int8   `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data"`
}
