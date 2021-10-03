package usecases

import (
	"fmt"
	"go_clean_api/api/adapters/gateways/repositories"
	"go_clean_api/api/entities"
)

type SignUpInteractor struct {
	repo repositories.UserRepository
}

func BuildSignUpInteractor(r repositories.UserRepository) *SignUpInteractor {
	return &SignUpInteractor{
		repo: r,
	}
}

func (bs *SignUpInteractor) Execute(user *entities.User) (*entities.User, error) {

	var err error

	fmt.Println("Building user entity")

	user, err = entities.BuildUser(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println("Search already user with email: ", user.Email)

	userExists, err := bs.repo.FindByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, fmt.Errorf("user already exists wiht email")
	}

	result, err := bs.repo.CreateUser(user)
	if err != nil {
		return result, err
	}

	fmt.Println("User created with success")

	return result, nil
}
