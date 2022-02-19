package vo

import (
	"errors"
	ERROR "go_clean_api/api/constants/errors"
)

type Name struct {
	Value string
}

func (n *Name) Validate() error {
	length := len(n.Value)

	if length <= 0 {
		return errors.New(ERROR.NAME_IS_EMPTY)
	}

	return nil
}
