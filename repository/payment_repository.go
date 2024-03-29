package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) CreateCallbackPayment(callback_payment model.Callback_payment) error {
	res := r.DB.Create(&callback_payment)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert callback payment")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllCallbackPayment() []model.Callback_payment {
	callback_payment := []model.Callback_payment{}
	r.DB.Find(&callback_payment)

	return callback_payment
}
