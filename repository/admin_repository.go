package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) GetAdminByUsername(username string) (admin model.Admin, err error) {
	res := r.DB.Where("username = ?", username).Find(&admin)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("admin not found")
	}

	return
}
