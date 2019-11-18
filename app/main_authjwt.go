package main

import (
	"AuthorizationJWT/adapter"
	"AuthorizationJWT/engine"
	"AuthorizationJWT/provider/fileconfig"

	"AuthorizationJWT/provider/postgres"
	"AuthorizationJWT/provider/redis"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// Load Configuration
	dbConfig := fileconfig.GetDBConfig()

	// Connect and Migrate DB
	db := postgres.NewStorage(dbConfig)
	// Connect Redis
	rds := redis.NewRedis()
	// Prepare Engine for Use Case Logic
	eng := engine.NewEngine(db, rds)
	// Start application
	adapter := adapter.Handler{}
	adapter.InitializeServer(eng)
	adapter.Run(":4985")

}
