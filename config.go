package elog

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Client   string                       `yaml:"client"`
	Settings map[string]map[string]string `yaml:"settings"`
}

func MustLoadConfig(path string) *Config {
	yamlfile, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlfile, &config)

	if err != nil {
		panic(err)
	}

	return &config
}

func LoadConfig(path string) (*Config, error) {
	yamlfile, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlfile, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
