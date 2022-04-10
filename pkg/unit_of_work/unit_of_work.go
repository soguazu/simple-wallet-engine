package tx

import (
	"gorm.io/gorm"
	"wallet_engine/internals/core/ports"
)

type gormUnitOfWork struct {
	db *gorm.DB
}

// NewGormUnitOfWork will create a new gorm unit of work
func NewGormUnitOfWork(db *gorm.DB) ports.IUnitOfWork {
	return &gormUnitOfWork{db: db}
}

func (u *gormUnitOfWork) Begin() (*gorm.DB, error) {
	tx := u.db.Begin()
	u.db = tx
	return tx, tx.Error
}

func (u *gormUnitOfWork) Commit() error {
	return u.db.Commit().Error
}

func (u *gormUnitOfWork) Rollback() error {
	return u.db.Rollback().Error
}
