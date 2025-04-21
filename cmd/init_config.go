package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gBarczyszyn/mitosis/internal/config"
	"github.com/spf13/cobra"
)

var initConfigCmd = &cobra.Command{
	Use:   "init-config",
	Short: "Create a default ~/.mitosis/config.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not determine home dir: %w", err)
		}

		configPath := filepath.Join(home, ".mitosis", "config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("✅ config.yaml already exists, skipping")
			return nil
		}

		if err := config.CreateDefaultConfig(configPath); err != nil {
			return fmt.Errorf("failed to create default config: %w", err)
		}

		fmt.Println("✅ Default config.yaml created at", configPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initConfigCmd)
}
