package handlers

type EncrypterHandler interface {
	EncryptPassword(p string) (string, error)
}
