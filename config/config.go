package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	database *database
}

var config *Config

func Database() *database {
	return config.database
}

func (c *Config) String() string {
	return fmt.Sprintf("[Database]: %v\n", c.database)
}

var ConfigMap *viper.Viper
var instance sync.Once

func LoadConfig() error {
	instance.Do(func() {
		v := viper.New()
		v.SetConfigName("application")
		v.AddConfigPath("./")
		v.AddConfigPath("../")
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err != nil {
			fmt.Printf("Failed read file config : %s", err.Error())
		}
		ConfigMap = v
	})

	database := database{
		host:     ConfigMap.GetString("db_host"),
		port:     ConfigMap.GetInt("db_port"),
		username: ConfigMap.GetString("db_username"),
		password: ConfigMap.GetString("db_password"),
		dbName:   ConfigMap.GetString("db_name"),
	}

	config = &Config{
		database: &database,
	}

	log.Printf("loading configuration: %s\n", config.String())
	return nil
}
