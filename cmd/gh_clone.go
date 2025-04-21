package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var ghCloneCmd = &cobra.Command{
	Use:   "clone [owner/repo]",
	Short: "Clone a GitHub repository into $HOME/github.com/owner/repo",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ownerRepo := args[0]
		parts := strings.Split(ownerRepo, "/")
		if len(parts) != 2 {
			return fmt.Errorf("invalid repository format. Expected owner/repo")
		}
		owner := parts[0]
		repo := parts[1]

		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get user home: %w", err)
		}

		targetDir := filepath.Join(homeDir, "github.com", owner, repo)
		if _, err := os.Stat(targetDir); err == nil {
			return fmt.Errorf("target directory already exists: %s", targetDir)
		}

		if err := os.MkdirAll(filepath.Dir(targetDir), 0755); err != nil {
			return fmt.Errorf("failed to create directory structure: %w", err)
		}

		repoURL := fmt.Sprintf("git@github.com:%s.git", ownerRepo)
		cmdGit := exec.Command("git", "clone", repoURL, targetDir)
		cmdGit.Stdout = os.Stdout
		cmdGit.Stderr = os.Stderr

		if err := cmdGit.Run(); err != nil {
			return fmt.Errorf("git clone failed: %w", err)
		}

		fmt.Printf("Repository cloned to %s\n", targetDir)
		return nil
	},
}

func init() {
	ghCmd.AddCommand(ghCloneCmd)
}
