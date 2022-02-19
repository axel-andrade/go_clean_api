package signup

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type SignUpGateway interface {
	CancelTransaction() error
	CreateUser(user *entities.User) (*entities.User, error)
	CommitTransaction() error
	EncryptPassword(p string) (string, error)
	FindUserByEmail(email string) (*entities.User, error)
	StartTransaction() error
	NextEntityID() string
}

type SignUpInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpOutputDTO struct {
	Data entities.User
}

type SignUpPresenter interface {
	Show(out *SignUpOutputDTO, err error) common.OutputPort
}
