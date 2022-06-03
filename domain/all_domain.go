package domain

import (
	"github.com/CapstoneProject31/backend_ppob_31/model"
)

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)
}

type AdapterService interface {
	LoginAdmin(username, password string) (string, int)
}
