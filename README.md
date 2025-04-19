# 🧬 mitosis

> Git-based workspace sync for UNIX devices, dotfiles-first.

**Mitosis** is a command-line tool to automatically sync your shell environment, dotfiles, and development setup across multiple Unix-based machines using Git.

## 🚀 Features

- ✅ Syncs system config files to a Git repo (e.g. `.zshrc`, `.gitconfig`, `~/.config/...`)
- ✅ Applies tracked files from Git into the system
- ✅ Watches for changes and auto-commits/pushes
- ✅ Simple YAML configuration
- ✅ Clean Code, no bloat, 100% Go
- ✅ Runs as a background systemd user service

## 🛠 Installation

### 1. One-line install (requires Go):

```bash
curl -sL https://raw.githubusercontent.com/gBarczyszyn/mitosis/main/install-with-service.sh | bash
```

This will:
- Clone the repository
- Build the binary
- Install it to `/usr/local/bin`
- Set up `mitosis` as a systemd user service

### 2. Create your `~/.mitosis/config.yaml`

```yaml
repo_url: git@github.com:gBarczyszyn/mitosis-dotfiles.git
tracked_paths:
  - ~/.zshrc
  - ~/.p10k.zsh
  - ~/.gitconfig
  - ~/.config/nvim/init.vim
```

### 3. Sync manually (optional)

```bash
mitosis sync --config ~/.mitosis/config.yaml
```

### 4. Apply from repo (optional)

```bash
mitosis apply --config ~/.mitosis/config.yaml
```

### 5. Run manually as daemon (optional)

```bash
mitosis daemon --config ~/.mitosis/config.yaml
```

## 📦 Directory structure

By default, the repo is cloned into:

```
~/.mitosis/<repo-name>/
```

## 🧑‍💻 License

MIT © 2025 [Guilherme Barczyszyn](https://github.com/gBarczyszyn)
