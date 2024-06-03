package handler

import (
	"net/http"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func InitProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

func (h *productHandler) GetByCategory(c *gin.Context) {
	param, err := contract.BuildGetListProductParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	resp, pagination, err := h.productService.GetByCategory(c.Request.Context(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, contract.APIResponse(resp, pagination))
}
