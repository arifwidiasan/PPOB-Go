package service

import (
	"github.com/CapstoneProject31/backend_ppob_31/config"
	"github.com/CapstoneProject31/backend_ppob_31/domain"
)

type svc struct {
	c    config.Config
	repo domain.AdapterRepository
}

func NewService(repo domain.AdapterRepository, c config.Config) domain.AdapterService {
	return &svc{
		repo: repo,
		c:    c,
	}
}
