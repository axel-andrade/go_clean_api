package repositories_impl

import (
	"fmt"
	"go_clean_api/api/entities"
)

type UserRepositoryImpl struct {
	BaseRepositoryImpl
}

func BuildUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{BaseRepositoryImpl: *BuildBaseRepoImpl()}
}

func (r *UserRepositoryImpl) CreateUser(user *entities.User) (*entities.User, error) {

	q := r.getQueryOrTx()

	err := q.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *entities.User) error {
	err := r.Db.Save(user).Error
	return err
}

func (r *UserRepositoryImpl) FindUserByEmail(email string) (*entities.User, error) {

	var user entities.User

	err := r.Db.Limit(1).Find(&user, "email = ?", email).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindUserByID(id entities.EntityID) (*entities.User, error) {

	var user entities.User
	r.Db.First(&user, "id = ?", id)

	if user.Token == "" {
		return nil, fmt.Errorf("user does not exist")
	}

	return &user, nil

}
