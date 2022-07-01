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
