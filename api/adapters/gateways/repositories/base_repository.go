package repositories

import "go_clean_api/api/entities"

type BaseRepository interface {
	StartTransaction() error
	CommitTransaction() error
	CancelTransaction() error
	NextEntityID() entities.UniqueEntityID
}
