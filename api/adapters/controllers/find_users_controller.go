package controllers

import (
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/find_users"
)

type FindUsersController struct {
	Interactor interactor.FindUsersInteractor
}

func (ctrl *FindUsersController) Run(input interactor.FindUserInputDTO) common.OutputPort {
	output := ctrl.Interactor.Execute(input)
	return output
}
