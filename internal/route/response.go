package route

type ServerResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}
