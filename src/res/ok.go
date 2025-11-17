package res

type Ok[T any] struct {
	Code int8 `json:"code"`
	Data T    `json:"data"`
}
