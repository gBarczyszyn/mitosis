package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type AWSConfig struct {
	Enabled            bool `yaml:"enabled"`
	IncludeCredentials bool `yaml:"include_credentials"`
}

type SSHConfig struct {
	Enabled bool     `yaml:"enabled"`
	Keys    []string `yaml:"keys"`
}

type GHConfig struct {
	Enabled      bool `yaml:"enabled"`
	IncludeHosts bool `yaml:"include_hosts"`
}

type NvimConfig struct {
	Enabled    bool   `yaml:"enabled"`
	ConfigPath string `yaml:"config_path"`
}

type VSCodeConfig struct {
	Enabled     bool `yaml:"enabled"`
	Settings    bool `yaml:"settings"`
	Keybindings bool `yaml:"keybindings"`
}

type Config struct {
	RepoURL      string       `yaml:"repo_url"`
	RepoPath     string       `yaml:"repo_path"`
	TrackedPaths []string     `yaml:"tracked_paths"`
	AWS          AWSConfig    `yaml:"aws"`
	SSH          SSHConfig    `yaml:"ssh"`
	GH           GHConfig     `yaml:"gh"`
	Nvim         NvimConfig   `yaml:"nvim"`
	VSCode       VSCodeConfig `yaml:"vscode"`
	Custom       []string     `yaml:"custom"`
}

func LoadConfig(userProvided string) (*Config, error) {
	var configPath string

	if userProvided != "" {
		configPath = userProvided
	} else {
		home, _ := os.UserHomeDir()
		mitosisPath := filepath.Join(home, ".mitosis")

		entries, err := os.ReadDir(mitosisPath)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ~/.mitosis: %v", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				testPath := filepath.Join(mitosisPath, entry.Name(), "config.yaml")
				if _, err := os.Stat(testPath); err == nil {
					configPath = testPath
					break
				}
			}
		}

		if configPath == "" {
			return nil, fmt.Errorf("no config.yaml found inside ~/.mitosis")
		}
	}

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	if cfg.RepoURL == "" {
		fmt.Println("⚠️  No repo_url configured — sync will be skipped.")
		return &cfg, nil
	}

	if cfg.RepoPath == "" {
		repoName := extractRepoName(cfg.RepoURL)
		homeDir, _ := os.UserHomeDir()
		cfg.RepoPath = filepath.Join(homeDir, ".mitosis", repoName)
	}

	paths := cfg.TrackedPaths

	if cfg.AWS.Enabled {
		paths = append(paths, "~/.aws/config")
		if cfg.AWS.IncludeCredentials {
			paths = append(paths, "~/.aws/credentials")
		}
	}

	if cfg.SSH.Enabled {
		for _, key := range cfg.SSH.Keys {
			paths = append(paths, filepath.Join("~/.ssh", key))
		}
	}

	if cfg.GH.Enabled {
		paths = append(paths, "~/.config/gh/config.yml")
		if cfg.GH.IncludeHosts {
			paths = append(paths, "~/.config/gh/hosts.yml")
		}
	}

	if cfg.Nvim.Enabled {
		paths = append(paths, cfg.Nvim.ConfigPath)
	}

	if cfg.VSCode.Enabled {
		if cfg.VSCode.Settings {
			paths = append(paths, "~/.config/Code/User/settings.json")
		}
		if cfg.VSCode.Keybindings {
			paths = append(paths, "~/.config/Code/User/keybindings.json")
		}
	}

	if len(cfg.Custom) > 0 {
		paths = append(paths, cfg.Custom...)
	}

	cfg.TrackedPaths = paths
	return &cfg, nil
}

func extractRepoName(url string) string {
	url = strings.TrimSuffix(url, ".git")
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func createDefaultConfig(path string) error {
	defaultCfg := Config{
		AWS:    AWSConfig{},
		SSH:    SSHConfig{Keys: []string{}},
		GH:     GHConfig{},
		Nvim:   NvimConfig{ConfigPath: "~/.config/nvim"},
		VSCode: VSCodeConfig{},
		Custom: []string{},
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	encoder.SetIndent(2)
	return encoder.Encode(defaultCfg)
}
