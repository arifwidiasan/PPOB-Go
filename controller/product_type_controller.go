package controller

import (
	"strconv"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// CreateProductType godoc
// @Summary Create New Product Type.
// @Description create new product type with name.
// @Tags Product Type
// @Accept json
// @Produce json
// @Param	product_type	body	docs.CreateProductType	true	"JSON name"
// @Success	201 {object} docs.CreateProductTypeSuccess
// @Failure 500 {object} docs.CreateProductTypeFail
// @Router /product_types [POST]
func (ce *EchoController) CreateProductTypeController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	product_type := model.Product_type{}
	if err := c.Bind(&product_type); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreateProductTypeService(product_type)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":          "success",
		"product_type_name": product_type.Name,
	})
}

func (ce *EchoController) GetAllProductTypeController(c echo.Context) error {

	product_types := ce.Svc.GetAllProductTypeService()

	return c.JSON(200, map[string]interface{}{
		"messages":      "success",
		"product_types": product_types,
	})
}

func (ce *EchoController) GetOneProductTypeController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetProductTypeByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "product type not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":     "success",
		"product_type": res,
	})
}

func (ce *EchoController) UpdateProductTypeController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	product_type := model.Product_type{}
	if err := c.Bind(&product_type); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.UpdateProductTypeByIDService(id_int, product_type)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeleteProductTypeController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.DeleteProductTypeByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "product type not found",
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}
