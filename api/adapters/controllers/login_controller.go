package controllers

import (
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/login"
)

type LoginController struct {
	Interactor interactor.LoginInteractor
}

func (ctrl *LoginController) Run(input interactor.LoginInputDTO) common.OutputPort {
	output := ctrl.Interactor.Execute(input)
	return output
}
