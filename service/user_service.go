package service

import (
	"fmt"

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
