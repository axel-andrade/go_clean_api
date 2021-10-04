package user

import (
	ERROR "go_clean_api/api/shared/constants"
	"go_clean_api/api/usecases/common"
	interactor "go_clean_api/api/usecases/user/signup"
)

type SignUpPresenter struct {
	Output common.OutputPort
}

func (p *SignUpPresenter) Show(out *interactor.SignUpOutputDTO, err error) common.OutputPort {

	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(out)
}

func (p *SignUpPresenter) formatSuccessOutput(out *interactor.SignUpOutputDTO) common.OutputPort {

	p.Output.Data = out
	p.Output.StatusCode = 201

	return p.Output
}

func (p *SignUpPresenter) formatErrOutput(err error) common.OutputPort {

	switch err.Error() {
	case ERROR.EMAIL_ALREADY_EXISTS:
		p.Output.StatusCode = 409
	case ERROR.NAME_IS_EMPTY, ERROR.INVALID_EMAIL, ERROR.INVALID_PASSWORD:
		p.Output.StatusCode = 422
	default:
		p.Output.StatusCode = 400
	}

	p.Output.Error = err.Error()

	return p.Output
}
