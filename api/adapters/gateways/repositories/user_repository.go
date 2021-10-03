package repositories

import (
	"go_clean_api/api/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(email string) (*entities.User, error)
}
