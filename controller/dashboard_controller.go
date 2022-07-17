package controller

import (
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) GetCountTransactionController(c echo.Context) error {

	count, _ := ce.Svc.GetCountTransactionService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"count":    count,
	})
}

func (ce *EchoController) GetCountTransactionSuccessController(c echo.Context) error {

	count, _ := ce.Svc.GetCountTransactionSuccessService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"count":    count,
	})
}

func (ce *EchoController) GetCountProductController(c echo.Context) error {

	count, _ := ce.Svc.GetCountProductService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"count":    count,
	})
}

func (ce *EchoController) GetCountUserController(c echo.Context) error {

	count, _ := ce.Svc.GetCountUserService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"count":    count,
	})
}
