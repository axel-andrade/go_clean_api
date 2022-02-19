package login

import (
	"fmt"
	ERROR "go_clean_api/api/constants/errors"
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type LoginInteractor struct {
	Gateway   LoginGateway
	Presenter LoginPresenter
}

func BuildLoginInteractor(g LoginGateway, p LoginPresenter) *LoginInteractor {
	return &LoginInteractor{Gateway: g, Presenter: p}
}

func (bs *LoginInteractor) Execute(input LoginInputDTO) common.OutputPort {

	fmt.Println("info: search already user with email: ", input.Email)
	user, err := bs.Gateway.FindUserByEmail(input.Email)

	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if user == nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERROR.USER_NOT_FOUND))
	}

	fmt.Println("info: comparing passwords")
	if err = bs.Gateway.CompareHashAndPassword(user.Password.Value, input.Password); err != nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERROR.INCORRECT_PASSWORD))
	}

	fmt.Println("info: generate token")
	td, err := bs.Gateway.GenerateToken(user.ID)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if err = bs.Gateway.CreateAuth(user.ID, td); err != nil {
		return bs.Presenter.Show(nil, err)
	}

	output := bs.formatOutput(user, td)

	return bs.Presenter.Show(&output, nil)
}

func (bs *LoginInteractor) formatOutput(user *entities.User, td *entities.TokenDetails) LoginOutputDTO {
	var out LoginOutputDTO

	out.User = *user
	out.AccessToken = td.AccessToken
	out.RefreshToken = td.RefreshToken

	return out
}
