package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall mitosis and clean up all configuration",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("⚠️  Are you sure you want to uninstall mitosis and delete all configs? (y/N): ")
		confirm, _ := reader.ReadString('\n')

		if confirm != "y\n" && confirm != "Y\n" {
			fmt.Println("❌ Aborted.")
			return
		}

		// Remove ~/.mitosis
		usr, err := user.Current()
		if err == nil {
			configPath := filepath.Join(usr.HomeDir, ".mitosis")
			os.RemoveAll(configPath)
			fmt.Println("🧹 Removed ~/.mitosis")
		}

		// Remove binary from /usr/local/bin
		binPath := "/usr/local/bin/mitosis"
		if err := os.Remove(binPath); err == nil {
			fmt.Println("🗑️  Removed", binPath)
		} else {
			fmt.Println("⚠️  Could not remove binary:", err)
		}

		fmt.Println("✅ Mitosis successfully uninstalled.")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
