package repositories

import "go_clean_api/api/entities"

type SessionRepository interface {
	GetAuth(auth *entities.AccessDetails) (string, error)
	CreateAuth(userid string, td *entities.TokenDetails) error
	DeleteAuth(uuid string) (int64, error)
}
