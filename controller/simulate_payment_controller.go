package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ce *EchoController) SimulatePaymentController(c echo.Context) error {
	code_transaction := c.Param("code_transaction")

	price := c.Param("price")
	price_int, _ := strconv.Atoi(price)

	err := ce.Svc.SimulatePaymentService(code_transaction, price_int)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
	})
}
