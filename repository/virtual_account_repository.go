package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) CreateVirtualAccount(virtual_account model.Virtual_account) error {
	res := r.DB.Create(&virtual_account)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert virtual account")
	}

	return nil
}

func (r *repositoryMysqlLayer) CheckVirtualAccount(user_id, payment_method_id int) (virtual_account model.Virtual_account, err error) {
	res := r.DB.Where("user_id = ? AND payment_method_id = ?", user_id, payment_method_id).
		Preload(clause.Associations).Find(&virtual_account)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("virtual account not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateVirtualAccountByID(id int, virtual_account model.Virtual_account) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&virtual_account)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update virtual account")
	}

	return nil
}
