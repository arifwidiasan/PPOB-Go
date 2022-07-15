package service

import (
	"fmt"

	"github.com/CapstoneProject31/backend_ppob_31/config"
	"github.com/CapstoneProject31/backend_ppob_31/domain"

	"github.com/golang-jwt/jwt"
)

type svc struct {
	c    config.Config
	repo domain.AdapterRepository
}

func (s *svc) ClaimToken(bearer *jwt.Token) string {
	claim := bearer.Claims.(jwt.MapClaims)
	username := fmt.Sprintf("%v", claim["username"])

	return username
}

func NewService(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svc{
		repo: repo,
		c:    c,
	}
}
