package controllers

import (
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/logout"
)

type LogoutController struct {
	Interactor interactor.LogoutInteractor
}

func (ctrl *LogoutController) Run(encodedToken string) common.OutputPort {
	output := ctrl.Interactor.Execute(encodedToken)
	return output
}
