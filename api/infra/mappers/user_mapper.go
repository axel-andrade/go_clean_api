package mappers

import (
	"go_clean_api/api/entities"
	vo "go_clean_api/api/entities/value_objects"
	"go_clean_api/api/infra/database/models"
)

type UserMapper struct {
	BaseMapper
}

func (m *UserMapper) ToDomain(model models.User) *entities.User {
	return &entities.User{
		Base:  *m.BaseMapper.toDomain(model.Base),
		Email: vo.Email{Value: model.Email},
		Name:  vo.Name{Value: model.Name},
	}
}

func (m *UserMapper) ToPersistence(entity entities.User) *models.User {
	return &models.User{
		Base:     *m.BaseMapper.toPersistence(entity.Base),
		Email:    entity.Email.Value,
		Name:     entity.Name.Value,
		Password: entity.Password.Value,
	}
}
