package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) CreateTransaction(transaction model.Transaction) error {
	res := r.DB.Create(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert transaction")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetTransactionByCodeTransaction(code_transaction string) (transaction model.Transaction, err error) {
	res := r.DB.Where("code_transaction = ?", code_transaction).Find(&transaction)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("code transaction not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateTransactionByID(id int, transaction model.Transaction) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update transaction")
	}

	return nil
}
