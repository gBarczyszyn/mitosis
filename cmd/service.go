package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the mitosis daemon",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "darwin" {
			exec.Command("launchctl", "load", filepath.Join(os.Getenv("HOME"), "Library/LaunchAgents/com.gbarczyszyn.mitosis.plist")).Run()
			fmt.Println("✅ Daemon started with launchctl")
		} else if runtime.GOOS == "linux" {
			exec.Command("systemctl", "--user", "start", "mitosis.service").Run()
			fmt.Println("✅ Daemon started with systemd")
		} else {
			fmt.Println("❌ Unsupported OS")
		}
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the mitosis daemon",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "darwin" {
			exec.Command("launchctl", "unload", filepath.Join(os.Getenv("HOME"), "Library/LaunchAgents/com.gbarczyszyn.mitosis.plist")).Run()
			fmt.Println("🛑 Daemon stopped with launchctl")
		} else if runtime.GOOS == "linux" {
			exec.Command("systemctl", "--user", "stop", "mitosis.service").Run()
			fmt.Println("🛑 Daemon stopped with systemd")
		} else {
			fmt.Println("❌ Unsupported OS")
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show daemon status",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "darwin" {
			exec.Command("launchctl", "list").Run()
			fmt.Println("ℹ️  Use 'launchctl list | grep mitosis' to see status")
		} else if runtime.GOOS == "linux" {
			exec.Command("systemctl", "--user", "status", "mitosis.service").Run()
		} else {
			fmt.Println("❌ Unsupported OS")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(statusCmd)
}
