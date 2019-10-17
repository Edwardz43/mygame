package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

// GetDBConfig ...
func GetDBConfig() string {
	dbHost := viper.GetString(`mysqldb.host`)
	dbPort := viper.GetString(`mysqldb.port`)
	dbUser := viper.GetString(`mysqldb.user`)
	dbPass := viper.GetString(`mysqldb.pass`)
	dbName := viper.GetString(`mysqldb.name`)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

// GetDBConfigV2 ...
func GetDBConfigV2() string {
	dbHost := viper.GetString(`postgresdb.host`)
	dbPort := viper.GetString(`postgresdb.port`)
	dbUser := viper.GetString(`postgresdb.user`)
	dbPass := viper.GetString(`postgresdb.pass`)
	dbName := viper.GetString(`postgresdb.name`)
	dbssl := viper.GetString(`postgresdb.sslmode`)
	// host=127.0.0.1 port=15432 user=admin dbname=postgres password=password sslmode=disable
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPass, dbssl)
}

// GetLogstashConfig ...
func GetLogstashConfig() string {
	host := viper.GetString(`logstash.host`)
	port := viper.GetString(`logstash.port`)
	return host + ":" + port
}

// GetELKConfig ...
func GetELKConfig() bool {
	isEnabled := viper.GetBool(`ELK.isEnabled`)
	return isEnabled
}

// GetTokenKey ...
func GetTokenKey() string {
	key := viper.GetString(`Token.key`)
	return key
}
