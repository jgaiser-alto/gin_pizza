package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Url string `yaml:"url"`
	} `yaml:"database"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/common/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
