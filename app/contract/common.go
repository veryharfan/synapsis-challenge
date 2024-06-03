package contract

import (
	"fmt"
	"math"
)

var (
	ErrInternalServer = fmt.Errorf("internal server error")
)

type Response struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func APIResponse[D any](data D, pagination *Pagination) Response {
	return Response{
		Success:    true,
		Data:       data,
		Pagination: pagination,
	}
}

func APIResponseErr(err error) ResponseError {
	return ResponseError{
		Success: false,
		Message: err.Error(),
	}
}

type Pagination struct {
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	TotalData int `json:"totalData"`
}

func GetPaginationData(page, limit, totalData int) *Pagination {
	totalPage := math.Ceil(float64(totalData) / float64(limit))

	return &Pagination{
		Page:      page,
		TotalPage: int(totalPage),
		TotalData: totalData,
	}
}
