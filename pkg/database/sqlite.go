package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"wallet_engine/internals/core/domain"
	"wallet_engine/internals/core/ports"
)

type sqliteDatastore struct {
}

// NewSqliteDatabase creates a new instance for managing database
func NewSqliteDatabase() ports.IDatabase {
	return &sqliteDatastore{}
}

func (d *sqliteDatastore) ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic("failed to connect database")
	}

	fmt.Println("Established database connection")

	d.MigrateAll(db)
	return db
}

func (d *sqliteDatastore) MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Wallet{})
}
