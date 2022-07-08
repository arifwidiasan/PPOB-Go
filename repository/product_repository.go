package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"gorm.io/gorm/clause"
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

func (r *repositoryMysqlLayer) GetAllProduct() []model.Product {
	products := []model.Product{}
	r.DB.Preload(clause.Associations).Find(&products)

	return products
}

func (r *repositoryMysqlLayer) GetProductByID(id int) (product model.Product, err error) {
	res := r.DB.Where("id = ?", id).Preload(clause.Associations).Find(&product)
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

func (r *repositoryMysqlLayer) GetProductByProductType(product_type_id int) (products []model.Product, err error) {
	res := r.DB.Where("product_type_id = ?", product_type_id).Preload(clause.Associations).
		Find(&products)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetProductByOperator(operator_id int) (products []model.Product, err error) {
	res := r.DB.Where("operator_id = ?", operator_id).Preload(clause.Associations).
		Find(&products)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}

	return
}
