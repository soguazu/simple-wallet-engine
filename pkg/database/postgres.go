package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"wallet_engine/internals/core/domain"
	"wallet_engine/internals/core/ports"
)

type datastore struct {
}

// NewDatabase creates a new instance for managing database
func NewDatabase() ports.IDatabase {
	return &datastore{}
}

func (d *datastore) ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic("failed to connect database")
	}

	fmt.Println("Established database connection")

	return db
}

func (d *datastore) MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Wallet{})
}
