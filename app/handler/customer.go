package handler

import (
	"errors"
	"net/http"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/service"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerService service.CustomerService
}

func InitCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{customerService}
}

func (h *customerHandler) Register(c *gin.Context) {
	var request contract.CustomerRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	err := h.customerService.Create(c.Request.Context(), request)
	if err != nil {
		if errors.Is(err, contract.ErrUsernameAlreadyExist) {
			c.JSON(http.StatusConflict, contract.APIResponseErr(contract.ErrUsernameAlreadyExist))
			return
		}
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, contract.APIResponse(map[string]interface{}{}))
}

func (h *customerHandler) Login(c *gin.Context) {
	var request contract.CustomerRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	resp, err := h.customerService.Login(c.Request.Context(), request)
	if err != nil {
		if errors.Is(err, contract.ErrWrongPassword) {
			c.JSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrWrongPassword))
			return
		}
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.IndentedJSON(http.StatusOK, contract.APIResponse(resp))
}
