package gitops

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func isGitRepo(dir string) bool {
	gitPath := filepath.Join(dir, ".git")
	info, err := os.Stat(gitPath)
	return err == nil && info.IsDir()
}

func pull(repoPath string) error {
	cmd := exec.Command("git", "rev-parse", "--verify", "main")
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		fmt.Println("⚠️  No 'main' branch yet — skipping pull")
		return nil
	}

	cmd = exec.Command("git", "pull", "--rebase")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func commitAndPush(repoPath, message string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "--no-verify", "-m", message)
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	cmd = exec.Command("git", "push")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func SyncWithPaths(repoURL, repoPath string, paths []string) error {
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		fmt.Println("🔄 Repository not found locally, cloning...")
		cmd := exec.Command("git", "clone", repoURL, repoPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to clone repository: %v", err)
		}
	}

	if !isGitRepo(repoPath) {
		return fmt.Errorf("directory %s is not a Git repository", repoPath)
	}

	if err := pull(repoPath); err != nil {
		return fmt.Errorf("git pull failed: %v", err)
	}

	for _, originalPath := range paths {
		expanded := originalPath
		if strings.HasPrefix(expanded, "~/") {
			home, _ := os.UserHomeDir()
			expanded = filepath.Join(home, strings.TrimPrefix(expanded, "~/"))
		}

		absSource, err := filepath.Abs(expanded)

		if err != nil {
			return fmt.Errorf("failed to resolve %s: %v", originalPath, err)
		}

		relPath := strings.TrimPrefix(expanded, os.Getenv("HOME"))
		relPath = strings.TrimPrefix(relPath, "/")

		dst := filepath.Join(repoPath, relPath)

		if _, err := os.Stat(absSource); os.IsNotExist(err) {
			fmt.Printf("⚠️  Skipping missing file: %s\n", absSource)
			continue
		}

		if err := copyFile(absSource, dst); err != nil {
			return fmt.Errorf("failed to copy %s to %s: %v", absSource, dst, err)
		}

	}

	msg := fmt.Sprintf("mitosis: auto sync at %s", time.Now().Format(time.RFC822))
	if err := commitAndPush(repoPath, msg); err != nil {
		return fmt.Errorf("commit/push failed: %v", err)
	}

	fmt.Println("✅ Sync complete")
	return nil
}
