package server

import (
	"gorm.io/gorm"
	"log"
	"wallet_engine/internals/core/ports"
	"wallet_engine/pkg/config"
)

// DBConnection stores the instance of the Database
var DBConnection *gorm.DB

// Run function starts the database connection
func Run(database ports.IDatabase) error {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
		return err
	}

	DBConnection = database.ConnectDB(config.Instance.DatabaseURL)
	err = database.MigrateAll(DBConnection)
	//_ = database.DropAll(DBConnection)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
