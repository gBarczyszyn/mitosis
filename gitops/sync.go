package gitops

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func isGitRepo(dir string) bool {
	gitPath := filepath.Join(dir, ".git")
	info, err := os.Stat(gitPath)
	return err == nil && info.IsDir()
}

func pull(repoPath string) error {
	cmd := exec.Command("git", "pull", "--rebase")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func commitAndPush(repoPath, message string) error {
	add := exec.Command("git", "add", ".")
	add.Dir = repoPath
	add.Stdout = os.Stdout
	add.Stderr = os.Stderr
	if err := add.Run(); err != nil {
		return err
	}

	commit := exec.Command("git", "commit", "-m", message)
	commit.Dir = repoPath
	commit.Stdout = os.Stdout
	commit.Stderr = os.Stderr
	_ = commit.Run()

	push := exec.Command("git", "push")
	push.Dir = repoPath
	push.Stdout = os.Stdout
	push.Stderr = os.Stderr
	return push.Run()
}

func Sync(repoPath string) error {
	fmt.Println("🧬 Starting Git sync:", repoPath)

	if !isGitRepo(repoPath) {
		return fmt.Errorf("directory %s is not a Git repository", repoPath)
	}

	if err := pull(repoPath); err != nil {
		return fmt.Errorf("git pull failed: %v", err)
	}

	msg := fmt.Sprintf("mitosis: auto sync at %s", time.Now().Format(time.RFC822))
	if err := commitAndPush(repoPath, msg); err != nil {
		return fmt.Errorf("commit/push failed: %v", err)
	}

	fmt.Println("✅ Git sync completed")
	return nil
}
