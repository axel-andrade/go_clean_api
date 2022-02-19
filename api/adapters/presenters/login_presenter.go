package presenters

import (
	ERROR "go_clean_api/api/constants/errors"
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/login"
	"net/http"
)

type LoginPresenter struct{}

func (p *LoginPresenter) Show(result *interactor.LoginOutputDTO, err error) common.OutputPort {

	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result)
}

func (p *LoginPresenter) formatSuccessOutput(result *interactor.LoginOutputDTO) common.OutputPort {

	var out common.OutputPort

	out.Data = result
	out.StatusCode = http.StatusOK

	return out
}

func (p *LoginPresenter) formatErrOutput(err error) common.OutputPort {

	var out common.OutputPort

	switch err.Error() {
	case ERROR.USER_NOT_FOUND:
		out.StatusCode = http.StatusNotFound
	case ERROR.INVALID_EMAIL, ERROR.INVALID_PASSWORD:
		out.StatusCode = http.StatusUnprocessableEntity
	default:
		out.StatusCode = http.StatusBadRequest
	}

	out.Error = err.Error()

	return out
}
