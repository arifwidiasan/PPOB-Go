package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) CreatePaymentMethod(payment_method model.Payment_method) error {
	res := r.DB.Create(&payment_method)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert payment method")
	}

	return nil
}
