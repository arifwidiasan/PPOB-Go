package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) CreateProductTypeService(product_type model.Product_type) error {
	if product_type.Name == "" {
		return fmt.Errorf("error insert product type")
	}
	return s.repo.CreateProductType(product_type)
}

func (s *svc) GetAllProductTypeService() []model.Product_type {
	return s.repo.GetAllProductType()
}

func (s *svc) GetProductTypeByIDService(id int) (model.Product_type, error) {
	return s.repo.GetProductTypeByID(id)
}

func (s *svc) UpdateProductTypeByIDService(id int, product_type model.Product_type) error {
	if product_type.Name == "" {
		return fmt.Errorf("error update product type")
	}
	return s.repo.UpdateProductTypeByID(id, product_type)
}

func (s *svc) DeleteProductTypeByIDService(id int) error {
	return s.repo.DeleteProductTypeByID(id)
}
