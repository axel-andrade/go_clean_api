package repositories_impl

import (
	"fmt"
	"go_clean_api/api/entities"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}

func (repo *UserRepoImpl) CreateUser(user *entities.User) (*entities.User, error) {

	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepoImpl) UpdateUser(user *entities.User) error {
	err := repo.Db.Save(user).Error
	return err
}

func (repo *UserRepoImpl) FindByEmail(email string) (*entities.User, error) {

	var user entities.User

	err := repo.Db.Limit(1).Find(&user, "email = ?", email).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepoImpl) FindByID(email string) (*entities.User, error) {

	var user entities.User
	repo.Db.First(&user, "id = ?", email)

	if user.Token == "" {
		return nil, fmt.Errorf("user does not exist")
	}

	return &user, nil

}
