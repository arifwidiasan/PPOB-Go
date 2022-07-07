package controller

import (
	"strconv"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	user, err := ce.Svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	transaction := model.Transaction{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	transaction.UserID = user.ID
	err = ce.Svc.CreateTransactionService(transaction)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages": "success",
	})
}

func (ce *EchoController) GetAllTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	transactions := ce.Svc.GetAllTransactionService()

	return c.JSON(200, map[string]interface{}{
		"messages":     "success",
		"transactions": transactions,
	})
}

func (ce *EchoController) GetOneTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetTransactionByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "transaction not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":    "success",
		"transaction": res,
	})
}

func (ce *EchoController) UpdateTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	transaction := model.Transaction{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.UpdateTransactionByIDService(id_int, transaction)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeleteTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.DeleteTransactionByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "transaction not found",
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoController) GetAllUserTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	user, err := ce.Svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	transactions := ce.Svc.GetAllUserTransactionService(int(user.ID))

	return c.JSON(200, map[string]interface{}{
		"messages":     "success",
		"transactions": transactions,
	})
}

func (ce *EchoController) GetOneUserTransactionController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	user, err := ce.Svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	code_transaction := c.Param("code_transaction")
	res, err := ce.Svc.GetTransactionByCodeTransactionService(code_transaction)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "transaction not found",
		})
	}

	if res.UserID != user.ID {
		return c.JSON(400, map[string]interface{}{
			"messages": "you don't have access to this transaction",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":    "success",
		"transaction": res,
	})
}
