package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `default:"denti"`
	Port    string `default:"8282"`
	Logger  struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"denti.log"`
	}
	DB struct {
		Use      string `default:"postgres"`
		Postgres struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"postgres"`
			Port     string `default:"5432"`
			UserName string `default:"postgres"`
			Password string `default:"postgres"`
			Database string `default:"kouyi_palm_treasure"`
		}
	}
	Contacts struct {
		Name  string `default:"Akbar Shaikh"`
		Email string `default:"aashaikh55@gmail.com"`
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	fmt.Print(cfg)
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
