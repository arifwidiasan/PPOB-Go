package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) GetCountTransaction() (count int64, err error) {
	res := r.DB.Model(&model.Transaction{}).Count(&count)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("transaction not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetCountTransactionSuccess() (count int64, err error) {
	res := r.DB.Model(&model.Transaction{}).Where("status = ?", "success").Count(&count)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("transaction not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetCountProduct() (count int64, err error) {
	res := r.DB.Model(&model.Product{}).Count(&count)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetCountUser() (count int64, err error) {
	res := r.DB.Model(&model.User{}).Count(&count)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}

	return
}
