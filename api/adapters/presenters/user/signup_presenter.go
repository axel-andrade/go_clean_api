package presenters

import (
	output "go_clean_api/api/adapters/presenters"
	"go_clean_api/api/entities"
	const_errors "go_clean_api/api/shared/constants"
)

type SignUpPresenter struct {
	Output output.OutputPort
}

func (p *SignUpPresenter) Show(u *entities.User, err error) output.OutputPort {

	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(u)
}

func (p *SignUpPresenter) formatSuccessOutput(u *entities.User) output.OutputPort {

	p.Output.Data = u
	p.Output.StatusCode = 201
	p.Output.Error = nil

	return p.Output
}
func (p *SignUpPresenter) formatErrOutput(err error) output.OutputPort {

	switch err.Error() {
	case const_errors.EMAIL_ALREADY_EXISTS:
		p.Output.StatusCode = 409
	default:
		p.Output.StatusCode = 400
	}

	p.Output.Error = err

	return p.Output
}
