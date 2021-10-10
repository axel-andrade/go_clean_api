package repositories

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type UserRepository interface {
	BaseRepository
	CreateUser(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
	FindUserByID(id entities.UniqueEntityID) (*entities.User, error)
	FindUsersPaginate(pagination *entities.PaginationOptions) (*common.PaginateOutput, error)
}
