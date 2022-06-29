package controller

import (
	"strconv"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateOperatorController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	operator := model.Operator{}
	if err := c.Bind(&operator); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreateOperatorService(operator)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":      "success",
		"operator_name": operator.Name,
	})
}

func (ce *EchoController) GetAllOperatorController(c echo.Context) error {

	operator := ce.Svc.GetAllOperatorService()

	return c.JSON(200, map[string]interface{}{
		"messages":  "success",
		"operators": operator,
	})
}

func (ce *EchoController) GetOneOperatorController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetOperatorByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "operator not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"operator": res,
	})
}

func (ce *EchoController) UpdateOperatorController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	operator := model.Operator{}
	if err := c.Bind(&operator); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.UpdateOperatorByIDService(id_int, operator)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeleteOperatorController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.DeleteOperatorByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "operator not found",
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}
