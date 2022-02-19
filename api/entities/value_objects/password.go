package vo

import (
	"errors"
	ERROR "go_clean_api/api/constants/errors"
)

type Password struct {
	Value string
}

func (p *Password) Validate() error {
	length := len(p.Value)

	if length >= 6 {
		return nil
	}

	return errors.New(ERROR.INVALID_PASSWORD)

}
