package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) CreatePaymentMethodService(payment_method model.Payment_method) error {
	if payment_method.CodeBank == "" {
		return fmt.Errorf("error insert payment method")
	}
	return s.repo.CreatePaymentMethod(payment_method)
}

func (s *svc) GetAllPaymentMethodService() []model.Payment_method {
	return s.repo.GetAllPaymentMethod()
}

func (s *svc) GetPaymentMethodByIDService(id int) (model.Payment_method, error) {
	return s.repo.GetPaymentMethodByID(id)
}

func (s *svc) UpdatePaymentMethodByIDService(id int, payment_method model.Payment_method) error {
	return s.repo.UpdatePaymentMethodByID(id, payment_method)
}

func (s *svc) DeletePaymentMethodByIDService(id int) error {
	return s.repo.DeletePaymentMethodByID(id)
}
