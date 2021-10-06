package repositories

import "go_clean_api/api/entities"

type TokenManagerRepository interface {
	GetAuth(auth *entities.AccessDetails) (string, error)
	CreateAuth(userid string, td *entities.TokenDetails) error
}
