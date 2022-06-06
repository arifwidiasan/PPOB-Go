package controller

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
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
	product_type := model.Product_type{}
	c.Bind(&product_type)

	err := ce.Svc.CreateProductTypeService(product_type)
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
