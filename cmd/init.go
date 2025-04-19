package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"
)

var repoURL string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new mitosis repository and config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if repoURL == "" {
			log.Fatal("Please provide a --repo URL")
		}

		repoName := extractRepoName(repoURL)
		home, _ := os.UserHomeDir()
		repoPath := filepath.Join(home, ".mitosis", repoName)

		if _, err := os.Stat(repoPath); err == nil {
			log.Fatalf("Repository already exists at %s", repoPath)
		}

		fmt.Println("🔗 Cloning repository...")
		cmdClone := exec.Command("git", "clone", repoURL, repoPath)
		cmdClone.Stdout = os.Stdout
		cmdClone.Stderr = os.Stderr
		cmdClone.Env = append(os.Environ(), "GIT_SSH_COMMAND=ssh")
		if err := cmdClone.Run(); err != nil {
			log.Fatalf("failed to clone repository: %v", err)
		}

		configFile := filepath.Join(repoPath, "config.yaml")
		f, err := os.Create(configFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		encoder := yaml.NewEncoder(f)
		encoder.SetIndent(2)
		encoder.Encode(map[string]string{
			"repo_url": repoURL,
		})

		fmt.Println("✅ Mitosis initialized successfully!")
	},
}

func init() {
	initCmd.Flags().StringVar(&repoURL, "repo", "", "Git repository URL")
	rootCmd.AddCommand(initCmd)
}

func extractRepoName(url string) string {
	url = filepath.Base(url)
	url = filepath.Base(url[:len(url)-len(filepath.Ext(url))])
	return url
}
