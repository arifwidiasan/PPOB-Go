package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) CreateCallbackPaymentService(callback_payment model.Callback_payment) error {
	if callback_payment.ExternalID == "" {
		return fmt.Errorf("error insert payment")
	}
	return s.repo.CreateCallbackPayment(callback_payment)
}

func (s *svc) GetAllCallbackPaymentService() []model.Callback_payment {
	return s.repo.GetAllCallbackPayment()
}
