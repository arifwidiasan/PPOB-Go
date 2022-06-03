package repository

import (
	"github.com/CapstoneProject31/backend_ppob_31/domain"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
