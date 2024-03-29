package service

import (
	"net/http"

	"github.com/CapstoneProject31/backend_ppob_31/helper"
	"github.com/CapstoneProject31/backend_ppob_31/model"
)

func (s *svc) LoginAdmin(username, password string) (string, int) {
	admin, err := s.repo.GetAdminByUsername(username)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	if admin.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateTokenAdmin(int(admin.ID), admin.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svc) GetAdminByUsernameService(username string) (model.Admin, error) {
	return s.repo.GetAdminByUsername(username)
}
