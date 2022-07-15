package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	nanoid "github.com/aidarkhanov/nanoid/v2"
)

func (s *svc) CreateProductService(product model.Product) error {
	if product.Name == "" {
		return fmt.Errorf("error insert product, product name is empty")
	}
	id, _ := nanoid.New()
	product.CodeProduct = id
	err := s.repo.CreateProduct(product)
	if err != nil {
		return err
	}

	res, err := s.repo.GetProductByName(product.Name)
	if err != nil {
		return err
	}

	res.CodeProduct = fmt.Sprintf("%s%02d%02d%04d", "P", res.ProductTypeID, res.OperatorID, int(res.ID))
	err = s.repo.UpdateProductByID(int(res.ID), res)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) GetAllProductService() []model.Product {
	return s.repo.GetAllProduct()
}

func (s *svc) GetProductByIDService(id int) (model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *svc) UpdateProductByIDService(id int, product model.Product) error {
	return s.repo.UpdateProductByID(id, product)
}

func (s *svc) DeleteProductByIDService(id int) error {
	return s.repo.DeleteProductByID(id)
}

func (s *svc) GetProductByProductTypeService(product_type_id int) ([]model.Product, error) {
	products, err := s.repo.GetProductByProductType(product_type_id)
	return products, err

}

func (s *svc) GetProductByOperatorService(operator_id int) ([]model.Product, error) {
	products, err := s.repo.GetProductByOperator(operator_id)
	return products, err

}
