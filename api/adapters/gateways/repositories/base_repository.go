package repositories

type BaseRepository interface {
	StartTransaction() error
	CommitTransaction() error
	CancelTransaction() error
}
