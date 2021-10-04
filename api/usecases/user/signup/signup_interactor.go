package signup

import (
	"fmt"
	"go_clean_api/api/adapters/gateways/handlers"
	"go_clean_api/api/adapters/gateways/repositories"
	"go_clean_api/api/entities"
	ERRO "go_clean_api/api/shared/constants"
	"go_clean_api/api/usecases/common"
)

type SignUpInteractor struct {
	Repo             repositories.UserRepository
	EncrypterHandler handlers.EncrypterHandler
	Presenter        SignUpPresenter
}

func BuildSignUpInteractor(r repositories.UserRepository, e handlers.EncrypterHandler, p SignUpPresenter) *SignUpInteractor {
	return &SignUpInteractor{
		Repo:             r,
		EncrypterHandler: e,
		Presenter:        p,
	}
}

func (bs *SignUpInteractor) Execute(i SignUpInputDTO) common.OutputPort {

	var err error

	fmt.Println("info: building user entity")

	u, err := entities.BuildUser(i.Name, i.Email, i.Password)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	err = bs.encryptPassword(u)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	fmt.Println("info: search already user with email: ", u.Email)

	userExists, err := bs.Repo.FindByEmail(u.Email)

	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if userExists != nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERRO.EMAIL_ALREADY_EXISTS))
	}

	result, err := bs.Repo.CreateUser(u)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	fmt.Println("info: user created with success")

	var out SignUpOutputDTO
	out.Data = *result

	return bs.Presenter.Show(&out, nil)
}

func (bs *SignUpInteractor) encryptPassword(u *entities.User) (err error) {

	fmt.Println("info: encrypting password")

	newp, err := bs.EncrypterHandler.EncryptPassword(u.Password)
	if err != nil {
		return fmt.Errorf("error during password encryption: %v", err)
	}

	u.Password = string(newp)

	return nil
}
