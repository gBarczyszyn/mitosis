package gitops

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ApplyToSystem(repoURL, repoPath string, paths []string) error {
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		fmt.Println("🔄 Repository not found locally, cloning...")
		cmd := exec.Command("git", "clone", repoURL, repoPath)
		cmd.Env = append(os.Environ(), "GIT_SSH_COMMAND=ssh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to clone repository: %v", err)
		}
	}

	for _, relPath := range paths {
		cleanRel := strings.TrimPrefix(relPath, "~/")
		src := filepath.Join(repoPath, cleanRel)
		dst := filepath.Join(os.Getenv("HOME"), cleanRel)

		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("failed to copy %s to %s: %v", src, dst, err)
		}
	}

	fmt.Println("✅ Apply complete")
	return nil
}
