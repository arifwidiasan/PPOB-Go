package service

import (
	"fmt"
	"net/http"

	"github.com/CapstoneProject31/backend_ppob_31/helper"
	"github.com/CapstoneProject31/backend_ppob_31/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *svc) CreateUserService(user model.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("error insert user")
	}

	pass := []byte(user.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hash)

	return s.repo.CreateUser(user)
}

func (s *svc) LoginUser(input, password string) (string, int) {
	user, err := s.repo.CheckLoginUser(input)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateTokenUser(int(user.ID), user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
