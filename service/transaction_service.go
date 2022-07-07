package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	nanoid "github.com/aidarkhanov/nanoid/v2"
)

func (s *svc) CreateTransactionService(transaction model.Transaction) error {
	if transaction.ChargeNumber == "" {
		return fmt.Errorf("error insert transaction, charge number is empty")
	}
	id, _ := nanoid.New()
	transaction.CodeTransaction = id

	product, err := s.repo.GetProductByID(int(transaction.ProductID))
	if err != nil {
		return err
	}

	_, err = s.repo.GetPaymentMethodByID(int(transaction.PaymentMethodID))
	if err != nil {
		return err
	}

	transaction.Price = product.Price
	err = s.repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	res, err := s.repo.GetTransactionByCodeTransaction(transaction.CodeTransaction)
	if err != nil {
		return err
	}

	res.CodeTransaction = fmt.Sprintf("%s%02d%02d%04d", "TR-", res.UserID, res.ProductID, int(res.ID))
	err = s.repo.UpdateTransactionByID(int(res.ID), res)
	if err != nil {
		return err
	}

	return nil
}
