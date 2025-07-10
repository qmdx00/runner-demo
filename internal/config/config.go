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

		Grid struct {
			Columns int `mapstructure:"columns"`
			Rows    int `mapstructure:"rows"`
		} `mapstructure:"grid"`

		Sprite struct {
			Width  int `mapstructure:"width"`
			Height int `mapstructure:"height"`
		} `mapstructure:"sprite"`

		Tile struct {
			Width  int `mapstructure:"width"`
			Height int `mapstructure:"height"`
		} `mapstructure:"tile"`
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
