package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Resource struct {
	Name           string
	Endpoint       string
	DestinationURL string
}

type AppConfig struct {
	Server struct {
		Host        string
		Listen_port string
	} `mapstructure:"server"`
	Resources []Resource `mapstructure:"resources"`
}

var Config *AppConfig

func LoadConfig(configPath string, configName string) (*AppConfig, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("skill issue loading config file: %w", err)
	}

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("skill issue parsing config file: %w", err)
	}

	Config = &appConfig
	return Config, nil
}
