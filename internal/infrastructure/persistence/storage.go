package persistence

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/amburskui/httpserver/internal/domain"
)

var ErrNotFound = errors.New("not found")

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Create(username, firstname, lastname, email, phone string) (*domain.User, error) {
	name, err := domain.NewUsername(username)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username:  name,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Phone:     phone,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Storage) Get(id domain.UserIdentity) (*domain.User, error) {
	user := &domain.User{ID: id}

	if err := s.db.First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *Storage) Update(id domain.UserIdentity, username, firstname, lastname, email, phone string) (*domain.User, error) {
	name, err := domain.NewUsername(username)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:        id,
		Username:  name,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Phone:     phone,
	}

	if err := s.db.Clauses(clause.Returning{}).Updates(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *Storage) Delete(id domain.UserIdentity) error {
	result := s.db.Delete(&domain.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
