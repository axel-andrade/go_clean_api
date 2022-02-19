package common_ptr

import (
	"go_clean_api/api/entities"
	"time"
)

type UserFormatted struct {
	ID        entities.UniqueEntityID `json:"id" bson:"id"`
	Email     string                  `json:"email" bson:"email"`
	Name      string                  `json:"name" bson:"name"`
	CreatedAt time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time               `json:"updated_at" bson:"updated_at"`
}

type UserPresenter struct{}

func (ptr *UserPresenter) Format(user entities.User) UserFormatted {
	return UserFormatted{
		ID:        user.ID,
		Email:     user.Email.Value,
		Name:      user.Name.Value,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
