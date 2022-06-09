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

func (s *svc) GetAllProductService() []model.ProductResponse {
	product_type := model.Product_type{}
	operator := model.Operator{}
	products := s.repo.GetAllProduct()
	for i, v := range products {
		product_type, _ = s.repo.GetProductTypeByID(int(v.ProductTypeID))
		operator, _ = s.repo.GetOperatorByID(int(v.OperatorID))
		products[i].ProductTypeName = product_type.Name
		products[i].OperatorName = operator.Name
	}
	return products
}

func (s *svc) GetProductByIDService(id int) (model.ProductResponse, error) {
	product, err := s.repo.GetProductByID(id)
	product_type, _ := s.repo.GetProductTypeByID(int(product.ProductTypeID))
	operator, _ := s.repo.GetOperatorByID(int(product.OperatorID))
	product.ProductTypeName = product_type.Name
	product.OperatorName = operator.Name
	return product, err
}

func (s *svc) UpdateProductByIDService(id int, product model.Product) error {
	return s.repo.UpdateProductByID(id, product)
}

func (s *svc) DeleteProductByIDService(id int) error {
	return s.repo.DeleteProductByID(id)
}
