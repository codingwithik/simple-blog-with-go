package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost            string `mapstructure:"POSTGRES_HOST"`
	DBUserName        string `mapstructure:"POSTGRES_USER"`
	DBUserPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	DBName            string `mapstructure:"POSTGRES_DB"`
	DBPort            string `mapstructure:"POSTGRES_PORT"`
	ServerPort        string `mapstructure:"PORT"`
	JWTSecret         string `mapstructure:"JWT_SECRET"`
	ExpiryTime        string `mapstructure:"TOKEN_EXP_TIME"`
	RefreshExpiryTime string `mapstructure:"REFRESH_TOKEN_EXP_TIME"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	//viper.SetConfigType("env")
	//viper.SetConfigName(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}

func GetConfig() *Config {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return &config
}
