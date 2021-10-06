package controllers

import (
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/signup"
)

type SignUpController struct {
	Interactor interactor.SignUpInteractor
}

func (ctrl *SignUpController) Run(input interactor.SignUpInputDTO) common.OutputPort {
	output := ctrl.Interactor.Execute(input)
	return output
}
