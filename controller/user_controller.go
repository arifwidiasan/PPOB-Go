package controller

import (
	"net/http"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateUserController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err := ce.Svc.CreateUserService(user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages": "success",
		"username": user.Username,
	})
}

func (ce *EchoController) LoginUserController(c echo.Context) error {
	userLogin := model.UserLogin{}

	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	token, statusCode := ce.Svc.LoginUser(userLogin.Input, userLogin.Password)
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "username atau password salah",
		}, "  ")

	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, "  ")
}
