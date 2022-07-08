package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	nanoid "github.com/aidarkhanov/nanoid/v2"
)

func (s *svc) CreateVirtualAccountService(transaction model.Transaction) error {
	virtual_account := model.Virtual_account{}

	id, _ := nanoid.New()

	virtual_account.VAID = id
	virtual_account.ExternalID = transaction.CodeTransaction

	number := transaction.Payment_Method.Range + uint64(transaction.User.ID)
	virtual_account.VANumber = fmt.Sprintf("%s%d", transaction.Payment_Method.CodeVA, number)

	virtual_account.Price = transaction.Price
	virtual_account.UserID = transaction.UserID
	virtual_account.PaymentMethodID = transaction.PaymentMethodID

	res, err := s.repo.CheckVirtualAccount(int(virtual_account.UserID), int(virtual_account.PaymentMethodID))
	if err != nil {
		err = s.repo.CreateVirtualAccount(virtual_account)
		if err != nil {
			return err
		}
	} else {
		res.ExternalID = virtual_account.ExternalID
		res.Price = virtual_account.Price
		err = s.repo.UpdateVirtualAccountByID(int(res.ID), res)
		if err != nil {
			return err
		}
	}

	return nil
}
