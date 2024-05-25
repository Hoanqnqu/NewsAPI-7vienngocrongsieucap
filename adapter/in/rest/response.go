package rest

type APIResponse[T any] struct {
	StatusCode int
	Message    string
	Data       T
}
type APIResponseLogin struct {
	StatusCode  int
	Message     string
	AccessToken string
}
