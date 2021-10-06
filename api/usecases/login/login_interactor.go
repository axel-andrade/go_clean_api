package login

import (
	"fmt"
	"go_clean_api/api/adapters/gateways/handlers"
	"go_clean_api/api/adapters/gateways/repositories"
	ERRO "go_clean_api/api/shared/constants"
	"go_clean_api/api/usecases/common"
)

type LoginInteractor struct {
	UserRepo         repositories.UserRepository
	TokenManagerRepo repositories.TokenManagerRepository
	Encrypter        handlers.EncrypterHandler
	TokenManager     handlers.TokenManagerHandler
	Presenter        LoginPresenter
}

func BuildLoginInteractor(
	ur repositories.UserRepository,
	tr repositories.TokenManagerRepository,
	e handlers.EncrypterHandler,
	t handlers.TokenManagerHandler,
	p LoginPresenter,
) *LoginInteractor {

	return &LoginInteractor{
		UserRepo:         ur,
		TokenManagerRepo: tr,
		Encrypter:        e,
		TokenManager:     t,
		Presenter:        p,
	}
}

func (bs *LoginInteractor) Execute(input LoginInputDTO) common.OutputPort {

	fmt.Println("info: search already user with email: ", input.Email)
	user, err := bs.UserRepo.FindByEmail(input.Email)

	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if user == nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERRO.USER_NOT_FOUND))
	}

	fmt.Println("info: comparing passwords")
	if err = bs.Encrypter.CompareHashAndPassword(user.Password, input.Password); err != nil {
		return bs.Presenter.Show(nil, fmt.Errorf(ERRO.INCORRECT_PASSWORD))
	}

	fmt.Println("info: generate token")
	td, err := bs.TokenManager.GenerateToken(user.ID)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	if err = bs.TokenManagerRepo.CreateAuth(user.ID, td); err != nil {
		return bs.Presenter.Show(nil, err)
	}

	var out LoginOutputDTO
	out.User = *user
	out.AccessToken = td.AccessToken
	out.RefreshToken = td.RefreshToken

	return bs.Presenter.Show(&out, nil)
}
