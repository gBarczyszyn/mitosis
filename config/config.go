package config

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RepoURL      string   `yaml:"repo_url"`
	RepoPath     string   `yaml:"repo_path"`
	TrackedPaths []string `yaml:"tracked_paths"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.RepoPath == "" {
		repoName := extractRepoName(cfg.RepoURL)
		homeDir, _ := os.UserHomeDir()
		cfg.RepoPath = filepath.Join(homeDir, ".mitosis", repoName)
	}

	return &cfg, nil
}

func extractRepoName(url string) string {
	url = strings.TrimSuffix(url, ".git")
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
