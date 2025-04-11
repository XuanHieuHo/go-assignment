package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	LocalServerAddress string   `mapstructure:"SERVER_ADDRESS"`
	DBConfig           DBConfig `mapstructure:"DBCONFIG"`
}

type DBConfig struct {
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// read config from file .env
// replace config from OS
// save OS
// unmarshal struct
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //.env

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Not found .env to read config")
		} else {
			return
		}
	}

	// support OS variable, not apply for .env variable
	viper.SetEnvPrefix("LOCAL")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	fmt.Println(config)
	return
}

func SetOrReplaceEnv(key, value string) error {
	// os.Setenv sẽ tự động thêm mới hoặc ghi đè nếu key đã tồn tại
	err := os.Setenv(key, value)
	if err != nil {
		return fmt.Errorf("failed to set env variable %s: %w", key, err)
	}

	return nil
}
