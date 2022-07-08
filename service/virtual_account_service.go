package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/config"
	"github.com/CapstoneProject31/backend_ppob_31/model"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func newTrue() *bool {
	b := true
	return &b
}

func (s *svc) CreateVirtualAccountService(transaction model.Transaction) error {
	virtual_account := model.Virtual_account{}

	VA_Number := transaction.Payment_Method.Range + uint64(transaction.ID)
	String_VA_Number := fmt.Sprintf("%d", VA_Number)

	xendit.Opt.SecretKey = config.InitConfiguration().XENDITH_API
	data := virtualaccount.CreateFixedVAParams{
		ExternalID:           transaction.CodeTransaction,
		BankCode:             transaction.Payment_Method.CodeBank,
		Name:                 transaction.User.Name,
		VirtualAccountNumber: String_VA_Number,
		IsClosed:             newTrue(),
		ExpectedAmount:       float64(transaction.Price),
	}

	resp, err2 := virtualaccount.CreateFixedVA(&data)
	if err2 != nil {
		return fmt.Errorf("error create VA on Xendith")
	}

	virtual_account.VANumber = fmt.Sprintf("%s%d", transaction.Payment_Method.CodeVA, VA_Number)
	virtual_account.VAID = resp.ID
	virtual_account.ExternalID = resp.ExternalID
	virtual_account.Price = uint(resp.ExpectedAmount)
	virtual_account.UserID = transaction.UserID
	virtual_account.PaymentMethodID = transaction.PaymentMethodID

	err := s.repo.CreateVirtualAccount(virtual_account)
	if err != nil {
		return err
	}

	return nil
}
