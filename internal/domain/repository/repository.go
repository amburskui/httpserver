package repository

import "github.com/amburskui/httpserver/internal/domain"

type UserRepository interface {
	Create(username, firstname, lastname, email, phone string) (*domain.User, error)
	Get(id domain.UserIdentity) (*domain.User, error)
	Update(id domain.UserIdentity, username, firstname, lastname, email, phone string) (*domain.User, error)
	Delete(id domain.UserIdentity) error
}
