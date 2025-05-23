package cmd

import (
	"log"

	"github.com/gBarczyszyn/mitosis/config"
	"github.com/gBarczyszyn/mitosis/gitops"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes files from the system into the Git repository",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		if err := gitops.SyncWithPaths(cfg.RepoURL, cfg.RepoPath, cfg.TrackedPaths); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
