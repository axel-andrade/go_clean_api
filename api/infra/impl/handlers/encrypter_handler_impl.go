package handlers_impl

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type EncrypterHandlerImpl struct{}

func (e *EncrypterHandlerImpl) EncryptPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return string(hash), fmt.Errorf("error during password encryption: %v", err)
	}

	return string(hash), nil
}

func (e *EncrypterHandlerImpl) CompareHashAndPassword(hash string, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err
}
