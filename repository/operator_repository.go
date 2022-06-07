package repository

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (r *repositoryMysqlLayer) CreateOperator(operator model.Operator) error {
	res := r.DB.Create(&operator)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert operator")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllOperator() []model.Operator {
	operator := []model.Operator{}
	r.DB.Find(&operator)

	return operator
}

func (r *repositoryMysqlLayer) GetOperatorByID(id int) (operator model.Operator, err error) {
	res := r.DB.Where("id = ?", id).Find(&operator)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("operator not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateOperatorByID(id int, operator model.Operator) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&operator)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update operator")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteOperatorByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Operator{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete operator")
	}

	return nil
}
