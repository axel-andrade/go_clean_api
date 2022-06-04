package repositories_impl

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/infra/database/models"
	"go_clean_api/api/infra/mappers"
	utils "go_clean_api/api/infra/utils"
)

type UserRepositoryImpl struct {
	BaseRepositoryImpl
	UserMapper mappers.UserMapper
}

func BuildUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{BaseRepositoryImpl: *BuildBaseRepoImpl()}
}

func (r *UserRepositoryImpl) CreateUser(user *entities.User) (*entities.User, error) {

	model := r.UserMapper.ToPersistence(*user)

	q := r.getQueryOrTx()

	err := q.Create(model).Error

	if err != nil {
		return nil, err
	}

	return r.UserMapper.ToDomain(*model), nil
}

func (r *UserRepositoryImpl) UpdateUser(user *entities.User) error {
	err := r.Db.Save(user).Error
	return err
}

func (r *UserRepositoryImpl) FindUserByEmail(email string) (*entities.User, error) {
	var user models.User

	err := r.Db.Limit(1).Find(&user, "email = ?", email).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return r.UserMapper.ToDomain(user), nil
}

func (r *UserRepositoryImpl) FindUserByID(id entities.UniqueEntityID) (*entities.User, error) {

	var user entities.User
	err := r.Db.First(&user, "id = ?", id).Error

	if err != nil || user.ID == "" {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindUsersPaginate(pagination *entities.PaginationOptions) (*entities.PaginateResult, error) {
	var users []*entities.User
	var result entities.PaginateResult

	r.Db.Scopes(utils.Paginate(r.Db, users, pagination, &result)).Find(&users)

	utils.FormatPaginateOutput(pagination, &result)
	result.Docs = users

	return &result, nil
}
