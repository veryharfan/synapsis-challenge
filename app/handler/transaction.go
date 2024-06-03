package handler

import (
	"net/http"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/service"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func InitTransactionHandler(transactionService service.TransactionService) transactionHandler {
	return transactionHandler{
		transactionService: transactionService,
	}
}

func (h *transactionHandler) Create(c *gin.Context) {
	var request contract.CheckoutRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	request.CustomerId = int64(c.GetFloat64("uid"))
	resp, err := h.transactionService.Checkout(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, contract.APIResponseErr(contract.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, contract.APIResponse(resp, nil))
}
