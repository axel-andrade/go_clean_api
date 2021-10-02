package usecases

import (
	"go_clean_api/api/adapters/gateways/repositories"
	"go_clean_api/api/entities"
	"log"
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
	err := user.Prepare()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	user, err = bs.repo.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}
