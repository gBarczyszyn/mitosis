# 📦 Changelog

## [Unreleased] - 2025-04-19

### ✨ Added
- `mitosis start`: start the background daemon via launchctl or systemd
- `mitosis stop`: stop the daemon
- `mitosis status`: check current status of daemon
- Interactive `install.sh` script with support for REPO_URL prompt
- Automatic addition of mitosis binary to PATH (if needed)
- `config.yaml` can now live inside the Git repo, no need for ~/.mitosis/config.yaml

### 🛠 Changed
- `install.sh` prompts the user to start the daemon after setup
- README now reflects full install + usage experience
- Improved `mitosis doctor` output

### 🧼 Fixed
- Git not committing tracked files when run as background service (launchctl env vars missing)
- Paths not being expanded correctly for dotfiles

---

_Made with ❤️ by [Guilherme Barczyszyn](https://github.com/gBarczyszyn)_
