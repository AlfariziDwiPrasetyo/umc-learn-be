package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path ...string) (*Config, error) {
	targetPath := "internal/configs/config.yaml"

	data, err := os.ReadFile(targetPath)
	if err != nil {
		log.Fatalf("failed to read config file (%s): %v", targetPath, err)
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse YAML config: %v", err)
		return nil, err
	}

	return &cfg, nil
}
