package presenters

import (
	ERROR "go_clean_api/api/constants/errors"
	"go_clean_api/api/usecases/common"
	"net/http"
)

type LogoutPresenter struct{}

func (p *LogoutPresenter) Show(err error) common.OutputPort {

	var out common.OutputPort

	if err == nil {
		out.StatusCode = http.StatusNoContent
		return out
	}

	switch err.Error() {
	case ERROR.UNAUTHORIZED:
		out.StatusCode = http.StatusUnauthorized
	default:
		out.StatusCode = http.StatusBadRequest
	}

	out.Error = err.Error()

	return out
}
