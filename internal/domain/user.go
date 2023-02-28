package domain

import "errors"

var ErrMaxUsernameLengthExceeded = errors.New("maximum field size exceeded")

const MaxUsernameLength = 256

type UserIdentity int64

type Username string

func NewUsername(name string) (Username, error) {
	if len(name) > MaxUsernameLength {
		return "", ErrMaxUsernameLengthExceeded
	}

	return Username(name), nil
}

type User struct {
	ID        UserIdentity `gorm:"id;type:bigint:primarykey;autoIncrement:true" json:"id"`
	Username  Username     `gorm:"column:username;type:text" json:"username"`
	FirstName string       `gorm:"column:firstname" json:"firstname"`
	LastName  string       `gorm:"column:lastname" json:"lastname"`
	Email     string       `gorm:"column:email" json:"email"`
	Phone     string       `gorm:"column:phone" json:"phone"`
}
