package login

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	User         entities.User `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `refresh_token`
}

type LoginPresenter interface {
	Show(out *LoginOutputDTO, err error) common.OutputPort
}
