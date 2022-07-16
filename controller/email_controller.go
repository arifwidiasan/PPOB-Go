package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ce *EchoController) SendEmailController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	transaction, err := ce.Svc.GetTransactionByIDService(id_int)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	virtual_account, err := ce.Svc.GetVirtualAccountService(transaction.CodeTransaction)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.SendEmailService(transaction.User.Email, transaction, virtual_account)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":         "success",
		"transaction_code": transaction.CodeTransaction,
	})
}
