package presenters

import (
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/find_users"
	"net/http"
)

type FindusersPresenter struct{}

func (p *FindusersPresenter) Show(result *interactor.FindUsersOutputDTO, err error) common.OutputPort {

	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result)
}

func (p *FindusersPresenter) formatSuccessOutput(result *interactor.FindUsersOutputDTO) common.OutputPort {

	var out common.OutputPort

	out.Data = result
	out.StatusCode = http.StatusOK

	return out
}

func (p *FindusersPresenter) formatErrOutput(err error) common.OutputPort {

	var out common.OutputPort

	out.StatusCode = http.StatusBadRequest
	out.Error = err.Error()

	return out
}
