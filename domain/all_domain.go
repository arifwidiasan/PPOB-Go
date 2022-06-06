package domain

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
)

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)

	CreateProductType(product_type model.Product_type) error
}

type AdapterService interface {
	LoginAdmin(username, password string) (string, int)

	CreateProductTypeService(product_type model.Product_type) error
}
