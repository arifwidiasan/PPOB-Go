package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) CreateCallbackPaymentService(callback_payment model.Callback_payment) error {
	if callback_payment.ExternalID == "" {
		return fmt.Errorf("error insert payment")
	}

	err := s.repo.CreateCallbackPayment(callback_payment)
	if err != nil {
		return err
	}

	transaction, err := s.repo.GetTransactionByCodeTransaction(callback_payment.ExternalID)
	if err != nil {
		return err
	}

	transaction.Status = "success"

	err = s.repo.UpdateTransactionByID(int(transaction.ID), transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) GetAllCallbackPaymentService() []model.Callback_payment {
	return s.repo.GetAllCallbackPayment()
}
