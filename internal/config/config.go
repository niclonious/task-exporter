package config

import (
	"errors"
	"fmt"
	"slices"

	"github.com/spf13/viper"
)

type Config struct {
	Port int    `mapstructure:"TASK_EX_PORT"`
	Env  string `mapstructure:"TASK_EX_ENV"`
}

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

// LoadConfig loads configuration from config.env file at path and from environment
func LoadConfig(path string) (Config, error) {
	viper.SetConfigFile(path + "/config.env")
	viper.AutomaticEnv()
	viper.SetDefault("TASK_EX_PORT", 8080)
	viper.SetDefault("TASK_EX_ENV", "dev")
	config := Config{}
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	if err := validateConfig(config); err != nil {
		return config, err
	}
	return config, nil
}

// validateConfig validates configuration values
func validateConfig(conf Config) (err error) {
	if conf.Port < 1024 || conf.Port > 65535 {
		err = errors.Join(err, fmt.Errorf("TASK_EX_PORT should be between 1024 and 65535, got %d", conf.Port))
	}
	if !slices.Contains([]string{EnvDev, EnvProd}, conf.Env) {
		err = errors.Join(err, fmt.Errorf("TASK_EX_ENV should be one of [%s, %s], got %s", EnvDev, EnvProd, conf.Env))
	}
	return
}
