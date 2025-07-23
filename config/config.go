package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port          string `yaml:"port"`
	AuthToken     string `yaml:"auth_token"`
	PredictAPIURL string `yaml:"predict_api_url"`
}

var Cfg Config

func LoadConfig() {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Ошибка чтения config.yml: %v", err)
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalf("Ошибка парсинга config.yml: %v", err)
	}
}
