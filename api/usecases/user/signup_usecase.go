package usecases

import (
	"fmt"
	"go_clean_api/api/adapters/gateways/repositories"
	output "go_clean_api/api/adapters/presenters"
	presenters "go_clean_api/api/adapters/presenters/user"
	"go_clean_api/api/entities"
	ERRO "go_clean_api/api/shared/constants"
)

type SignUpInteractor struct {
	Repo      repositories.UserRepository
	Presenter presenters.SignUpPresenter
}

func BuildSignUpInteractor(r repositories.UserRepository) *SignUpInteractor {
	return &SignUpInteractor{
		Repo: r,
	}
}

func (bs *SignUpInteractor) Execute(user *entities.User) output.OutputPort {

	var err error

	fmt.Println("Building user entity")

	user, err = entities.BuildUser(user.Name, user.Email, user.Password)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	fmt.Println("Search already user with email: ", user.Email)

	userExists, err := bs.Repo.FindByEmail(user.Email)

	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if userExists != nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERRO.EMAIL_ALREADY_EXISTS))
	}

	result, err := bs.Repo.CreateUser(user)
	if err != nil {
		return bs.Presenter.Show(result, err)
	}

	fmt.Println("User created with success")

	return bs.Presenter.Show(result, nil)
}
