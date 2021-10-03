package usecases

import (
	"fmt"
	"go_clean_api/api/adapters/gateways/handlers"
	"go_clean_api/api/adapters/gateways/repositories"
	output "go_clean_api/api/adapters/presenters"
	presenters "go_clean_api/api/adapters/presenters/user"
	"go_clean_api/api/entities"
	ERRO "go_clean_api/api/shared/constants"
)

type SignUpInteractor struct {
	Repo             repositories.UserRepository
	EncrypterHandler handlers.EncrypterHandler
	Presenter        presenters.SignUpPresenter
}

func BuildSignUpInteractor(r repositories.UserRepository, e handlers.EncrypterHandler, p presenters.SignUpPresenter) *SignUpInteractor {
	return &SignUpInteractor{
		Repo:             r,
		EncrypterHandler: e,
		Presenter:        p,
	}
}

func (i *SignUpInteractor) Execute(u *entities.User) output.OutputPort {

	var err error

	err = i.encryptPassword(u)
	if err != nil {
		return i.Presenter.Show(nil, err)
	}

	fmt.Println("info: building user entity")

	u, err = entities.BuildUser(u.Name, u.Email, u.Password)
	if err != nil {
		return i.Presenter.Show(nil, err)
	}

	fmt.Println("info: search already user with email: ", u.Email)

	userExists, err := i.Repo.FindByEmail(u.Email)

	if err != nil {
		return i.Presenter.Show(nil, err)
	}

	if userExists != nil {
		return i.Presenter.Show(nil, fmt.Errorf(ERRO.EMAIL_ALREADY_EXISTS))
	}

	result, err := i.Repo.CreateUser(u)
	if err != nil {
		return i.Presenter.Show(result, err)
	}

	fmt.Println("info: user created with success")

	return i.Presenter.Show(result, nil)
}

func (i *SignUpInteractor) encryptPassword(u *entities.User) (err error) {

	fmt.Println("info: encrypting password")

	newp, err := i.EncrypterHandler.EncryptPassword(u.Password)
	if err != nil {
		return fmt.Errorf("error during password encryption: %v", err)
	}

	u.Password = string(newp)

	return nil
}
