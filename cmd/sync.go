package cmd

import (
	"log"

	"github.com/gBarczyszyn/mitosis/gitops"
	"github.com/spf13/cobra"
)

var repoPath string

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes the local repository using Git",
	Run: func(cmd *cobra.Command, args []string) {
		if repoPath == "" {
			log.Fatal("repository path is required (use --repo flag)")
		}

		if err := gitops.Sync(repoPath); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	syncCmd.Flags().StringVar(&repoPath, "repo", "", "Path to the dotfiles Git repository")
	rootCmd.AddCommand(syncCmd)
}
