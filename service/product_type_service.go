package service

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) CreateProductTypeService(product_type model.Product_type) error {
	return s.repo.CreateProductType(product_type)
}
