package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceHost string `yaml:"service_host"`
	ServicePort int    `yaml:"service_port"`
	Minio       `yaml:"minio"`
}

type Minio struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Endpoint string `yaml:"endpoint"`
}

func NewConfig(log *logrus.Logger) (*Config, error) {
	var err error

	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	log.Info("config parsed")
	log.Info(cfg.ServiceHost)
	log.Info(cfg.ServicePort)
	log.Info(cfg.Minio)

	return cfg, nil
}
