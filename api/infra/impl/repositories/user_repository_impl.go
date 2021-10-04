package repositories_impl

import (
	"fmt"
	"go_clean_api/api/entities"
)

type UserRepoImpl struct {
	Base BaseRepositoryImpl
}

func BuildUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{Base: *BuildBaseRepoImpl()}
}

func (r *UserRepoImpl) CreateUser(user *entities.User) (*entities.User, error) {

	q := r.Base.getQueryOrTx()

	err := q.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImpl) UpdateUser(user *entities.User) error {
	err := r.Base.Db.Save(user).Error
	return err
}

func (r *UserRepoImpl) FindByEmail(email string) (*entities.User, error) {

	var user entities.User

	err := r.Base.Db.Limit(1).Find(&user, "email = ?", email).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoImpl) FindByID(email string) (*entities.User, error) {

	var user entities.User
	r.Base.Db.First(&user, "id = ?", email)

	if user.Token == "" {
		return nil, fmt.Errorf("user does not exist")
	}

	return &user, nil

}

func (r *UserRepoImpl) StartTransaction() error {
	return r.Base.StartTransaction()
}

func (r *UserRepoImpl) CancelTransaction() error {
	return r.Base.CancelTransaction()
}

func (r *UserRepoImpl) CommitTransaction() error {
	return r.Base.CommitTransaction()
}
