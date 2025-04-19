package cmd

import (
	"log"

	"github.com/gBarczyszyn/mitosis/config"
	"github.com/gBarczyszyn/mitosis/gitops"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Applies tracked files from the repository into the system",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		if err := gitops.ApplyToSystem(cfg.RepoURL, cfg.RepoPath, cfg.TrackedPaths); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
