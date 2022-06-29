package controller

import (
	"strconv"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateProductController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreateProductService(product)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":     "success",
		"product_name": product.Name,
	})
}

func (ce *EchoController) GetAllProductController(c echo.Context) error {

	products := ce.Svc.GetAllProductService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"products": products,
	})
}

func (ce *EchoController) GetOneProductController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetProductByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "product not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"product":  res,
	})
}

func (ce *EchoController) UpdateProductController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.UpdateProductByIDService(id_int, product)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeleteProductController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.DeleteProductByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "product not found",
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}
