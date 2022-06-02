package controller

import (
	"net/http"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) LoginAdminController(c echo.Context) error {
	adminLogin := model.AdminLogin{}

	c.Bind(&adminLogin)

	token, statusCode := ce.Svc.LoginAdmin(adminLogin.Username, adminLogin.Password)
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
