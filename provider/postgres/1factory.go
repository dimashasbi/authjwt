package postgres

import (
	"AccountManagement/engine"
	"AccountManagement/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	storageFactory struct {
		db *gorm.DB
	}
)

// NewStorage creates a new instance of this mongodb storage factory
func NewStorage(data model.DBConfigurationModel) engine.StorageFactory {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Host, data.Port, data.User, data.Password, data.Dbname)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect Database " + err.Error())
	}
	fmt.Println("Database Connect")

	// migrate table and colomn (setting increment, null, size)
	MigrateResult := db.AutoMigrate(&model.Users{})
	if MigrateResult.Error != nil {
		panic("failed  Migrate Database " + MigrateResult.Error.Error())
	}
	fmt.Println("Database Migrate Success")

	// input index
	dbsetIndex(db)

	// done
	fmt.Println("Database Init Success")
	return &storageFactory{db}

}

func (c *storageFactory) NewUsersRespository() engine.UsersRepository {
	return newUsersRepository(c.db)
}

// dbsetIndex to Set Index
func dbsetIndex(db *gorm.DB) {

}
