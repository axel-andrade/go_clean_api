package repositories

import "go_clean_api/api/entities"

type SessionRepository interface {
	GetAuth(auth *entities.AccessDetails) (entities.EntityID, error)
	CreateAuth(userid entities.EntityID, td *entities.TokenDetails) error
	DeleteAuth(uuid string) (int64, error)
}
