package controller

import (
	"net/http"
	"strconv"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
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

func (ce *EchoController) GetAllUserController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	users := ce.Svc.GetAllUserService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}

func (ce *EchoController) GetOneUserController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetUserByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "user not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"user":     res,
	})
}

func (ce *EchoController) UpdateUserController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.UpdateUserByIDService(id_int, user)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeleteUserController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.DeleteUserByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "user not found",
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}
