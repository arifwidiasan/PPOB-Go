package controller

import (
	"strconv"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreatePaymentMethodController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	payment_method := model.Payment_method{}
	if err := c.Bind(&payment_method); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreatePaymentMethodService(payment_method)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":  "success",
		"bank_name": payment_method.Name,
		"bank_code": payment_method.CodeBank,
	})
}

func (ce *EchoController) GetAllPaymentMethodController(c echo.Context) error {

	payment_methods := ce.Svc.GetAllPaymentMethodService()

	return c.JSON(200, map[string]interface{}{
		"messages":        "success",
		"payment_methods": payment_methods,
	})
}

func (ce *EchoController) GetOnePaymentMethodController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetPaymentMethodByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "payment method not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":       "success",
		"payment_method": res,
	})
}

func (ce *EchoController) UpdatePaymentMethodController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	payment_method := model.Payment_method{}
	if err := c.Bind(&payment_method); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.UpdatePaymentMethodByIDService(id_int, payment_method)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeletePaymentMethodController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.DeletePaymentMethodByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "payment method not found",
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}
