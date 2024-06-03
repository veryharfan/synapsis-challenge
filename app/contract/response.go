package contract

import "fmt"

var (
	ErrInternalServer = fmt.Errorf("internal server error")
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func APIResponse[D any](data D) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func APIResponseErr(err error) ResponseError {
	return ResponseError{
		Success: false,
		Message: err.Error(),
	}
}
