package repositories

import (
	"go_clean_api/api/entities"
)

type UserRepository interface {
	Insert(user *entities.User) (*entities.User, error)
	Find(email string) (*entities.User, error)
}
