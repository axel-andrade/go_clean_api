package login

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type LoginGateway interface {
	CreateAuth(userid entities.EntityID, td *entities.TokenDetails) error
	CompareHashAndPassword(hash string, p string) error
	FindUserByEmail(email string) (*entities.User, error)
	GenerateToken(userid entities.EntityID) (*entities.TokenDetails, error)
}
type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginOutputDTO struct {
	User         entities.User `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

type LoginPresenter interface {
	Show(out *LoginOutputDTO, err error) common.OutputPort
}
