package handler

import (
	"net/http"
	"strconv"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/service"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	cartService service.CartService
}

func InitCartHandler(cartService service.CartService) cartHandler {
	return cartHandler{
		cartService: cartService,
	}
}

func (h *cartHandler) AddCart(c *gin.Context) {
	var request contract.CartRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	request.CustomerId = int64(c.GetFloat64("uid"))

	err := h.cartService.AddToCart(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, contract.APIResponse(gin.H{}, nil))
}

func (h *cartHandler) GetByCustomerId(c *gin.Context) {
	customerId := int64(c.GetFloat64("uid"))

	resp, err := h.cartService.GetByCustomerId(c.Request.Context(), customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, contract.APIResponse(resp, nil))
}

func (h *cartHandler) DeleteByCustomerIdAndProductId(c *gin.Context) {
	customerId := int64(c.GetFloat64("uid"))
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, contract.APIResponseErr(contract.ErrBadRequest))
		return
	}

	err = h.cartService.DeleteByCustomerIdAndProductId(c.Request.Context(), customerId, int64(productId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, contract.APIResponse(gin.H{}, nil))
}
