package signup

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

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
