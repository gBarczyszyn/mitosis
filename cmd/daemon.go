package cmd

import (
	"log"

	"github.com/gBarczyszyn/mitosis/config"
	"github.com/gBarczyszyn/mitosis/watcher"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Runs Mitosis in watch mode, syncing on change",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		watcher.StartWatcher(cfg)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
