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

	err = s.CreateVirtualAccountService(res)
	if err != nil {
		_ = s.repo.DeleteTransactionByID(int(res.ID))
		return err
	}

	return nil
}

func (s *svc) GetAllTransactionService() []model.Transaction {
	return s.repo.GetAllTransaction()
}

func (s *svc) GetTransactionByIDService(id int) (model.Transaction, error) {
	return s.repo.GetTransactionByID(id)
}

func (s *svc) UpdateTransactionByIDService(id int, transaction model.Transaction) error {
	return s.repo.UpdateTransactionByID(id, transaction)
}

func (s *svc) DeleteTransactionByIDService(id int) error {
	return s.repo.DeleteTransactionByID(id)
}

func (s *svc) GetAllUserTransactionService(id int) []model.Transaction {
	return s.repo.GetAllUserTransaction(id)
}

func (s *svc) GetTransactionByCodeTransactionService(code_transaction string) (model.Transaction, error) {
	return s.repo.GetTransactionByCodeTransaction(code_transaction)
}
