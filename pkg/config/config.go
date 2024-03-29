package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ChartConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type Config struct {
	Charts []ChartConfig `yaml:"charts"`
}

// LoadConfig reads a YAML configuration file and unmarshals it into a Config struct.
func LoadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	return config, err
}
