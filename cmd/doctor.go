package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gBarczyszyn/mitosis/config"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Prints resolved configuration and tracked paths",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		fmt.Println("🔍 Resolved Mitosis configuration:")

		if cfg.RepoURL == "" {
			fmt.Println("• Repo URL:  ⚠️ not configured")
		} else {
			fmt.Printf("• Repo URL:  %s\n", cfg.RepoURL)
		}

		fmt.Printf("• Repo Path: %s\n", cfg.RepoPath)

		gitDir := filepath.Join(cfg.RepoPath, ".git")
		if _, err := os.Stat(gitDir); os.IsNotExist(err) {
			fmt.Println("• Git Repo:   ❌ not cloned")
		} else {
			fmt.Println("• Git Repo:   ✅ cloned")
		}

		fmt.Println("• Tracked Paths:")
		for _, path := range cfg.TrackedPaths {
			expanded := os.ExpandEnv(path)
			if strings.HasPrefix(expanded, "~/") {
				home, _ := os.UserHomeDir()
				expanded = filepath.Join(home, expanded[2:])
			}

			if _, err := os.Stat(expanded); err == nil {
				fmt.Printf("  - %s ✅\n", path)
			} else {
				fmt.Printf("  - %s ❌ not found\n", path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
