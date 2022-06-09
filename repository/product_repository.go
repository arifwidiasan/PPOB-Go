package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) CreateProduct(product model.Product) error {
	res := r.DB.Create(&product)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert product")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetProductByName(name string) (product model.Product, err error) {
	res := r.DB.Where("name = ?", name).Find(&product)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateProductByID(id int, product model.Product) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&product)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update product")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllProduct() []model.ProductResponse {
	products := []model.ProductResponse{}
	r.DB.Model(&model.Product{}).
		//Select("products.*, product_types.name, operators.name").
		//Joins("JOIN product_types on product_types.id = products.product_type_id").
		//Joins("JOIN operators on operators.id = products.operator_id").
		Scan(&products)

	return products
}

func (r *repositoryMysqlLayer) GetProductByID(id int) (product model.ProductResponse, err error) {
	res := r.DB.Model(&model.Product{}).Where("id = ?", id).
		//Select("products.*, product_types.name, operators.name").
		//Joins("JOIN product_types on product_types.id = products.product_type_id").
		//Joins("JOIN operators on operators.id = products.operator_id").
		Scan(&product)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}

	return
}

func (r *repositoryMysqlLayer) DeleteProductByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Product{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete product")
	}

	return nil
}
