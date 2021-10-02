package repositories_impl

import (
	"fmt"
	"go_clean_api/api/entities"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func (repo UserRepoImpl) Insert(user *entities.User) (*entities.User, error) {

	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepoImpl) Find(email string) (*entities.User, error) {

	var user entities.User
	repo.Db.First(&user, "email = ?", email)

	if user.Token == "" {
		return nil, fmt.Errorf("User does not exist")
	}

	return &user, nil

}
