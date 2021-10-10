package handlers

type EncrypterHandler interface {
	EncryptPassword(p string) (string, error)
	CompareHashAndPassword(hash string, p string) error
}
