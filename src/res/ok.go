package res

type Ok[T any] struct {
	Code uint8 `json:"code"`
	Data T     `json:"data"`
}
