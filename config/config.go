package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Host string `json:"host"`
	PORT string `json:"port"`
}

var config *Config

func init() {
	basePath, _ := os.Getwd()

	configPath := filepath.Join(basePath, "config", "config.json")

	byt, _ := ioutil.ReadFile(configPath)

	_ = json.Unmarshal(byt, &config)

	if config.Host == "" {
		config.Host = "0.0.0.0"
	}

	if config.PORT == "" {
		config.PORT = "8000"
	}
}

func GetConfig() Config {
	return *config
}
