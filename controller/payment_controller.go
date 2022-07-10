package controller

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateCallbackPaymentController(c echo.Context) error {
	callback_payment := model.Callback_payment{}
	if err := c.Bind(&callback_payment); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err := ce.Svc.CreateCallbackPaymentService(callback_payment)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":         "success make payment",
		"callback_payment": callback_payment.ExternalID,
	})
}

func (ce *EchoController) GetAllCallbackPaymentController(c echo.Context) error {
	callback_payment := ce.Svc.GetAllCallbackPaymentService()

	return c.JSON(200, map[string]interface{}{
		"messages":         "success",
		"callback_payment": callback_payment,
	})
}
