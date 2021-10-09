package repositories

import "go_clean_api/api/entities"

type SessionRepository interface {
	GetAuth(auth *entities.AccessDetails) (entities.UniqueEntityID, error)
	CreateAuth(userid entities.UniqueEntityID, td *entities.TokenDetails) error
	DeleteAuth(uuid string) (int64, error)
}
