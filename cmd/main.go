package main

import (
	"log"
	"wallet_engine/cmd/server"
	_ "wallet_engine/docs"
	"wallet_engine/pkg/database"
)

// @title Wallet Engine API
// @version 1.0
// @description Wallet Engine API.
// @termsOfService http://swagger.io/terms/

// @contact.name Team API Support
// @contact.email info@test.com

// @license.name MIT
// @license.url https://github.com/sguazu

// @BasePath /v1
func main() {
	var DBConnection = database.NewDatabase()
	err := server.Run(DBConnection)
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection()
}
