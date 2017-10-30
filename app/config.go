package app

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	address  string
	mode     string
	logLevel string
	db       dbConfig
}

type dbConfig struct {
	driver  string
	source  string
	logMode bool
}

func loadConfig() config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	viper.SetDefault("address", ":8080")
	viper.SetDefault("mode", "release")
	viper.SetDefault("logLevel", "info")
	viper.SetDefault("database.driver", "mysql")
	viper.SetDefault("database.source", "root:root@/bookshelf?charset=utf8&parseTime=True&loc=Local")
	viper.SetDefault("database.logMode", false)

	address := viper.GetString("address")
	mode := viper.GetString("mode")
	logLevel := viper.GetString("logLevel")
	dbDriver := viper.GetString("database.driver")
	dbSource := viper.GetString("database.source")
	dbLogMode := viper.GetBool("database.logMode")

	return config{
		address:  address,
		mode:     mode,
		logLevel: logLevel,
		db: dbConfig{
			driver:  dbDriver,
			source:  dbSource,
			logMode: dbLogMode,
		},
	}
}
