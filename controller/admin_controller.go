package controller

import (
	"net/http"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/labstack/echo/v4"
)

// LoginAdmin godoc
// @Summary Login Admin.
// @Description login admin to get jwt token.
// @Tags Admin
// @Accept json
// @Produce json
// @Param	admin	body	docs.NewAdminLogin	true	"JSON username and user_pass"
// @Success	200	{object} docs.LoginAdminSuccess
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /admin/login [POST]
func (ce *EchoController) LoginAdminController(c echo.Context) error {
	adminLogin := model.AdminLogin{}

	if err := c.Bind(&adminLogin); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

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
