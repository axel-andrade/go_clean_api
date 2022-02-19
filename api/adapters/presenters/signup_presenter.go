package presenters

import (
	common_ptr "go_clean_api/api/adapters/presenters/common"
	ERROR "go_clean_api/api/constants/errors"
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/signup"
	"net/http"
)

type SignUpPresenter struct {
	UserPtr common_ptr.UserPresenter
}

func (p *SignUpPresenter) Show(result *interactor.SignUpOutputDTO, err error) common.OutputPort {

	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result)
}

func (p *SignUpPresenter) formatSuccessOutput(result *interactor.SignUpOutputDTO) common.OutputPort {

	var out common.OutputPort

	out.Data = p.UserPtr.Format(result.Data)
	out.StatusCode = http.StatusCreated

	return out
}

func (p *SignUpPresenter) formatErrOutput(err error) common.OutputPort {

	var out common.OutputPort

	switch err.Error() {
	case ERROR.EMAIL_ALREADY_EXISTS:
		out.StatusCode = http.StatusConflict
	case ERROR.NAME_IS_EMPTY, ERROR.INVALID_EMAIL, ERROR.INVALID_PASSWORD:
		out.StatusCode = http.StatusUnprocessableEntity
	default:
		out.StatusCode = http.StatusBadRequest
	}

	out.Error = err.Error()

	return out
}
