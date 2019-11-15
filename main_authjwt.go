package main

import (
	"AccountManagement/adapter"
	"AccountManagement/app/fileconfig"
	"AccountManagement/engine"

	"AccountManagement/provider/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// Load Configuration
	dbConfig := fileconfig.GetDBConfig()

	// Connect and Migrate DB
	db := postgres.NewStorage(dbConfig)
	// Prepare Engine for Use Case Logic
	eng := engine.NewEngine(db)
	// Start application
	adapter := adapter.Handler{}
	adapter.InitializeServer(eng)
	adapter.Run(":4983")

}
