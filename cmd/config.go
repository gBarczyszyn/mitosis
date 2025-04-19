package cmd

import (
	"fmt"
	"os"
	"strings"

	"mitosis/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage mitosis configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current repository URL",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadRepoConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("📦 Current repository URL:")
		fmt.Println(cfg.RepoURL)
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <repo_url>",
	Short: "Set repository URL in ~/.mitosis/repo.yaml",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := strings.TrimSpace(args[0])
		if err := config.SaveRepoConfig(repoURL); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			os.Exit(1)
		}
		fmt.Println("✅ Repository URL updated")
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	rootCmd.AddCommand(configCmd)
}
