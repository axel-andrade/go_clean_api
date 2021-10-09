package signup

import (
	"fmt"
	"go_clean_api/api/entities"
	ERRO "go_clean_api/api/shared/constants"
	"go_clean_api/api/usecases/common"
)

type SignUpInteractor struct {
	Gateway   SignUpGateway
	Presenter SignUpPresenter
}

func BuildSignUpInteractor(g SignUpGateway, p SignUpPresenter) *SignUpInteractor {
	return &SignUpInteractor{Gateway: g, Presenter: p}
}

func (bs *SignUpInteractor) Execute(i SignUpInputDTO) common.OutputPort {

	fmt.Println("info: building user entity")

	u, err := entities.BuildUser(i.Name, i.Email, i.Password)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if err = bs.encryptPassword(u); err != nil {
		return bs.Presenter.Show(nil, err)
	}

	fmt.Println("info: search already user with email: ", u.Email)

	userExists, err := bs.Gateway.FindUserByEmail(u.Email)

	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if userExists != nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERRO.EMAIL_ALREADY_EXISTS))
	}

	bs.Gateway.StartTransaction()

	result, err := bs.Gateway.CreateUser(u)
	if err != nil {
		bs.Gateway.CancelTransaction()
		return bs.Presenter.Show(nil, err)
	}

	bs.Gateway.CommitTransaction()

	fmt.Println("info: user created with success")

	var out SignUpOutputDTO
	out.Data = *result

	return bs.Presenter.Show(&out, nil)
}

func (bs *SignUpInteractor) encryptPassword(u *entities.User) (err error) {

	fmt.Println("info: encrypting password")

	newp, err := bs.Gateway.EncryptPassword(u.Password)
	if err != nil {
		return fmt.Errorf("error during password encryption: %v", err)
	}

	u.Password = string(newp)

	return nil
}
