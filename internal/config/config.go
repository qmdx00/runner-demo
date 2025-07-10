package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Global *Config
)

type Config struct {
	Game struct {
		Title string  `mapstructure:"title"`
		Scale float64 `mapstructure:"scale"`

		Window struct {
			Width  int `mapstructure:"width"`
			Height int `mapstructure:"height"`
		} `mapstructure:"window"`

		Cell struct {
			Width  int `mapstructure:"width"`
			Height int `mapstructure:"height"`
		} `mapstructure:"cell"`
	} `mapstructure:"game"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err = viper.Unmarshal(&Global); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}
}
