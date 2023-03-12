package entity

type ResponseError struct {
	Error string `json:"error"`
}

func NewResponseError(err string) ResponseError {
	return ResponseError{err}
}
