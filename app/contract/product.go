package contract

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetListProductParam struct {
	Page     int
	Limit    int
	Offset   int
	Sort     string
	Category string
}

type ProductResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildGetListProductParam(c *gin.Context) (param GetListProductParam, err error) {
	page, limit := 1, 10
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")
	sort := c.Query("sort")
	category := c.Query("category")

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			return
		}
	}

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			return
		}
	}

	offset := (page - 1) * limit

	param = GetListProductParam{
		Page:     page,
		Limit:    limit,
		Offset:   offset,
		Sort:     sort,
		Category: category,
	}
	return
}
