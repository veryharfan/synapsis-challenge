package handler

import (
	"errors"
	"net/http"
	"synapsis-challenge/src/contract"
	"synapsis-challenge/src/helper"
	"synapsis-challenge/src/service"

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
		c.JSON(http.StatusInternalServerError, helper.APIResponseErr(helper.ErrInternalServer))
		return
	}

	err := h.customerService.Create(c.Request.Context(), request)
	if err != nil {
		if errors.Is(err, contract.ErrUsernameAlreadyExist) {
			c.JSON(http.StatusConflict, helper.APIResponseErr(contract.ErrUsernameAlreadyExist))
			return
		}
		c.JSON(http.StatusInternalServerError, helper.APIResponseErr(helper.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse(map[string]interface{}{}))
}

func (h *customerHandler) Login(c *gin.Context) {
	var request contract.CustomerRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponseErr(helper.ErrInternalServer))
		return
	}

	resp, err := h.customerService.Login(c.Request.Context(), request)
	if err != nil {
		if errors.Is(err, contract.ErrWrongPassword) {
			c.JSON(http.StatusUnauthorized, helper.APIResponseErr(contract.ErrWrongPassword))
			return
		}
		c.JSON(http.StatusInternalServerError, helper.APIResponseErr(helper.ErrInternalServer))
		return
	}

	c.IndentedJSON(http.StatusOK, helper.APIResponse(resp))
}
