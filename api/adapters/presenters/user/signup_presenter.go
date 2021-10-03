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

	var out output.OutputPort

	if err != nil {
		return formatErrOutput(out, err)
	}

	out.Data = u
	out.StatusCode = 201

	return out
}

func formatErrOutput(out output.OutputPort, err error) output.OutputPort {

	switch err.Error() {
	case const_errors.EMAIL_ALREADY_EXISTS:
		out.StatusCode = 409
	default:
		out.StatusCode = 400
	}

	out.Error = err.Error()

	return out
}
