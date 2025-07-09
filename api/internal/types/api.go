package types

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Message string `json:"msg"`
}
