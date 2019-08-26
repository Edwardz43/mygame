package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// GetDBConfig ...
func GetDBConfig() string {
	// viper.SetConfigFile(`config.json`)
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
