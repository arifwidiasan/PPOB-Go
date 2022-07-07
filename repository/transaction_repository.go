package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) CreateTransaction(transaction model.Transaction) error {
	res := r.DB.Create(&transaction)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert transaction")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetTransactionByCodeTransaction(code_transaction string) (transaction model.Transaction, err error) {
	res := r.DB.Where("code_transaction = ?", code_transaction).Preload(clause.Associations).Find(&transaction)
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

func (r *repositoryMysqlLayer) GetAllTransaction() []model.Transaction {
	transactions := []model.Transaction{}
	r.DB.Preload(clause.Associations).Find(&transactions)

	return transactions
}

func (r *repositoryMysqlLayer) GetTransactionByID(id int) (transaction model.Transaction, err error) {
	res := r.DB.Where("id = ?", id).Preload(clause.Associations).Find(&transaction)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("transaction not found")
	}

	return
}

func (r *repositoryMysqlLayer) DeleteTransactionByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Transaction{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete transaction")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllUserTransaction(id int) []model.Transaction {
	transactions := []model.Transaction{}
	r.DB.Where("user_id = ?", id).Preload(clause.Associations).Find(&transactions)

	return transactions
}
