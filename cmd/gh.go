package cmd

import (
	"github.com/spf13/cobra"
)

var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "GitHub-related commands",
}

func init() {
	rootCmd.AddCommand(ghCmd)
}
