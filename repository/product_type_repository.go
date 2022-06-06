package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) CreateProductType(product_type model.Product_type) error {
	res := r.DB.Create(&product_type)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert product type")
	}

	return nil
}
