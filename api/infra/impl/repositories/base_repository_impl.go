package repositories_impl

import (
	"go_clean_api/api/entities"
	database "go_clean_api/api/infra/database"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseRepositoryImpl struct {
	Db *gorm.DB
	Tx *gorm.DB
}

func BuildBaseRepoImpl() *BaseRepositoryImpl {
	db := database.GetDB()
	return &BaseRepositoryImpl{Db: db, Tx: nil}
}

func (r *BaseRepositoryImpl) getQueryOrTx() *gorm.DB {
	if r.Tx != nil {
		return r.Tx
	}

	return r.Db
}

func (r *BaseRepositoryImpl) StartTransaction() error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := r.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	r.Tx = tx

	return nil
}

func (r *BaseRepositoryImpl) CommitTransaction() error {
	err := r.Tx.Commit().Error
	r.Tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl) CancelTransaction() error {
	err := r.Tx.Rollback().Error
	r.Tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepositoryImpl) NextEntityID() entities.UniqueEntityID {
	return uuid.NewV4().String()
}
