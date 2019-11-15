package fileconfig

import (
	"AccountManagement/model"
	"strconv"

	"github.com/spf13/viper"
)

// GetDBConfig to start get file
func GetDBConfig() model.DBConfigurationModel {
	// mylog.MnDebug("Load DB Configuration ")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("appconfig")

	dbhost := model.DBConfigurationModel{}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// put to Model
	dbhost.Host = viper.GetString("db.host")
	port, err := strconv.Atoi(viper.GetString("db.port"))
	dbhost.Port = port
	dbhost.User = viper.GetString("db.user")
	dbhost.Password = viper.GetString("db.password")
	dbhost.Dbname = viper.GetString("db.dbname")
	return dbhost
}
