package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	Redis RedisConfig
}

type AppConfig struct {
	Port     int    `mapstructure:"port"`
	LogLevel string `mapstructure:"log_level"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func LoadConfig() (Config, error) {
	var config Config

	// Set the file name and type
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Set the path where viper should look for the config file
	viper.AddConfigPath("./configs")

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	// Optionally: Override with environment variables
	viper.AutomaticEnv()

	return config, nil
}
