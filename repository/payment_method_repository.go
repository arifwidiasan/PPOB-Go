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

func (r *repositoryMysqlLayer) GetAllPaymentMethod() []model.Payment_method {
	payment_methods := []model.Payment_method{}
	r.DB.Find(&payment_methods)

	return payment_methods
}

func (r *repositoryMysqlLayer) GetPaymentMethodByID(id int) (payment_method model.Payment_method, err error) {
	res := r.DB.Where("id = ?", id).Find(&payment_method)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("payment method not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdatePaymentMethodByID(id int, payment_method model.Payment_method) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&payment_method)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update payment method")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeletePaymentMethodByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Payment_method{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete payment method")
	}

	return nil
}
