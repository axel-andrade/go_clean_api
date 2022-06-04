package signup

import (
	"fmt"
	ERROR "go_clean_api/api/constants/errors"
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type SignUpInteractor struct {
	Gateway   SignUpGateway
	Presenter SignUpPresenter
}

func BuildSignUpInteractor(g SignUpGateway, p SignUpPresenter) *SignUpInteractor {
	return &SignUpInteractor{Gateway: g, Presenter: p}
}

func (bs *SignUpInteractor) Execute(input SignUpInputDTO) common.OutputPort {
	fmt.Println("info: building user entity")

	nextId := bs.Gateway.NextEntityID()

	user, err := entities.BuildUser(input.Name, input.Email, input.Password, nextId)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if err = bs.encryptPassword(user); err != nil {
		return bs.Presenter.Show(nil, err)
	}

	fmt.Println("info: search already user with email: ", user.Email)

	userExists, err := bs.Gateway.FindUserByEmail(user.Email.Value)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if userExists != nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERROR.EMAIL_ALREADY_EXISTS))
	}

	bs.Gateway.StartTransaction()

	result, err := bs.Gateway.CreateUser(user)
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

	newp, err := bs.Gateway.EncryptPassword(u.Password.Value)
	if err != nil {
		return fmt.Errorf("error during password encryption: %v", err)
	}

	u.Password.Value = string(newp)

	return nil
}
