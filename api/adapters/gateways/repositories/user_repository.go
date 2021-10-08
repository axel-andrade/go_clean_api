package repositories

import (
	"go_clean_api/api/entities"
)

type UserRepository interface {
	BaseRepository
	CreateUser(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(id entities.EntityID) (*entities.User, error)
}
