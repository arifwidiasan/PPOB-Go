package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) CreateOperatorService(operator model.Operator) error {
	if operator.Name == "" {
		return fmt.Errorf("error insert operator")
	}
	return s.repo.CreateOperator(operator)
}

func (s *svc) GetAllOperatorService() []model.Operator {
	return s.repo.GetAllOperator()
}

func (s *svc) GetOperatorByIDService(id int) (model.Operator, error) {
	return s.repo.GetOperatorByID(id)
}

func (s *svc) UpdateOperatorByIDService(id int, operator model.Operator) error {
	if operator.Name == "" {
		return fmt.Errorf("error update operator")
	}
	return s.repo.UpdateOperatorByID(id, operator)
}

func (s *svc) DeleteOperatorByIDService(id int) error {
	return s.repo.DeleteOperatorByID(id)
}
