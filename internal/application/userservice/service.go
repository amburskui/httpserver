package userservice

import (
	"github.com/amburskui/httpserver/internal/domain"
	"github.com/amburskui/httpserver/internal/domain/repository"
)

type Service struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(username, firstname, lastname, email, phone string) (*domain.User, error) {
	return s.repo.Create(username, firstname, lastname, email, phone)
}

func (s *Service) Get(id domain.UserIdentity) (*domain.User, error) {
	return s.repo.Get(id)
}

func (s *Service) Update(id domain.UserIdentity, username, firstname, lastname, email, phone string) (*domain.User, error) {
	return s.repo.Update(id, username, firstname, lastname, email, phone)
}

func (s *Service) Delete(id domain.UserIdentity) error {
	return s.repo.Delete(id)
}
