package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func ImportConfig(path string, useEnv bool) (Config, error) {
	c := DefaultConfig()

	if path != "" {
		f, err := os.Open(path)
		if err != nil {
			return Config{}, fmt.Errorf("open config file: %w", err)
		}

		defer f.Close()

		err = yaml.NewDecoder(f).Decode(&c)
		if err != nil {
			return Config{}, fmt.Errorf("decode yaml: %w", err)
		}
	}

	if useEnv {
		// Важно: не понятно баг или фича, но nil значения инициализируются даже если нет окружения.
		err := envconfig.Process("APP", &c)
		if err != nil {
			return Config{}, fmt.Errorf("decode env: %w", err)
		}
	}

	return c, nil
}
