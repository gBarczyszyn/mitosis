package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gBarczyszyn/mitosis/config"
	"github.com/gBarczyszyn/mitosis/gitops"
)

func StartWatcher(cfg *config.Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	for _, path := range cfg.TrackedPaths {
		expanded := expandHome(path)
		fmt.Println("👁️  Watching:", expanded)
		if err := watcher.Add(expanded); err != nil {
			log.Fatalf("Failed to watch %s: %v", expanded, err)
		}
	}

	changeDetected := make(chan struct{}, 1)

	go func() {
		var timer *time.Timer

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) != 0 {
					fmt.Println("🔁 Change detected:", event.Name)
					if timer != nil {
						timer.Stop()
					}
					timer = time.AfterFunc(5*time.Second, func() {
						changeDetected <- struct{}{}
					})
				}
			case err := <-watcher.Errors:
				fmt.Println("⚠️  Watcher error:", err)
			}
		}
	}()

	for {
		select {
		case <-changeDetected:
			fmt.Println("🚀 Syncing changes...")
			err := gitops.SyncWithPaths(cfg.RepoURL, cfg.RepoPath, cfg.TrackedPaths)
			if err != nil {
				fmt.Println("❌ Sync failed:", err)
			}
		}
	}
}

func expandHome(path string) string {
	if len(path) > 2 && path[:2] == "~/" {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}
