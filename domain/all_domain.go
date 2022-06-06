package domain

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/golang-jwt/jwt"
)

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)

	CreateProductType(product_type model.Product_type) error
	GetAllProductType() []model.Product_type
	GetProductTypeByID(id int) (product_type model.Product_type, err error)
	UpdateProductTypeByID(id int, product_type model.Product_type) error
	DeleteProductTypeByID(id int) error
}

type AdapterService interface {
	LoginAdmin(username, password string) (string, int)
	GetAdminByUsernameService(username string) (model.Admin, error)

	ClaimToken(bearer *jwt.Token) string

	CreateProductTypeService(product_type model.Product_type) error
	GetAllProductTypeService() []model.Product_type
	GetProductTypeByIDService(id int) (model.Product_type, error)
	UpdateProductTypeByIDService(id int, product_type model.Product_type) error
	DeleteProductTypeByIDService(id int) error
}
