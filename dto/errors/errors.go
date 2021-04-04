package errors

type ErrorResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
}
