package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type RepoConfig struct {
	RepoURL string `yaml:"repo_url"`
}

func LoadRepoConfig() (*RepoConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to detect user home: %w", err)
	}

	path := filepath.Join(home, ".mitosis", "repo.yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("❌ repo.yaml not found. Run `mitosis install` or create ~/.mitosis/repo.yaml manually")
		}
		return nil, fmt.Errorf("failed to read repo.yaml: %w", err)
	}

	var cfg RepoConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.New("❌ invalid YAML format in repo.yaml")
	}

	if cfg.RepoURL == "" {
		return nil, errors.New("❌ repo.yaml is missing `repo_url`. Please define it")
	}

	return &cfg, nil
}

func SaveRepoConfig(url string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	configDir := filepath.Join(home, ".mitosis")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	path := filepath.Join(configDir, "repo.yaml")
	data := []byte(fmt.Sprintf("repo_url: %s\n", url))

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write repo.yaml: %w", err)
	}

	return nil
}
