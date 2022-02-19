package vo

import (
	"errors"
	ERROR "go_clean_api/api/constants/errors"
	"regexp"
)

type Email struct {
	Value string
}

func (e *Email) Validate() error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if regex.MatchString(e.Value) {
		return nil
	}

	return errors.New(ERROR.INVALID_EMAIL)
}
