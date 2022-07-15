package service

import (
	"github.com/CapstoneProject31/backend_ppob_31/helper"
	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) SendEmailService(email string, transaction model.Transaction, virtual_account model.Virtual_account) error {
	err := helper.SendEmail(email, transaction, virtual_account)
	if err != nil {
		return err
	}
	return nil
}
