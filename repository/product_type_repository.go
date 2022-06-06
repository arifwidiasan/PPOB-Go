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

func (r *repositoryMysqlLayer) GetAllProductType() []model.Product_type {
	product_types := []model.Product_type{}
	r.DB.Find(&product_types)

	return product_types
}

func (r *repositoryMysqlLayer) GetProductTypeByID(id int) (product_type model.Product_type, err error) {
	res := r.DB.Where("id = ?", id).Find(&product_type)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("product type not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateProductTypeByID(id int, product_type model.Product_type) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&product_type)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update product type")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteProductTypeByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Product_type{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete product type")
	}

	return nil
}
