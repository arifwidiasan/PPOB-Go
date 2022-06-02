package service

import (
	"net/http"

	"github.com/CapstoneProject31/backend_ppob_31/helper"
)

func (s *svc) LoginAdmin(username, password string) (string, int) {
	admin, _ := s.repo.GetAdminByUsername(username)

	if admin.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateTokenAdmin(int(admin.ID), admin.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
