# 🧬 mitosis

> Git-based workspace sync for UNIX devices, dotfiles-first.

**Mitosis** is a command-line tool to automatically sync your shell environment, dotfiles, and development setup across multiple Unix-based machines using Git.

## 🚀 Features

- ✅ Syncs system config files to a Git repo (e.g. `.zshrc`, `.gitconfig`, `~/.config/...`)
- ✅ Applies tracked files from Git into the system
- ✅ Watches for changes and auto-commits/pushes
- ✅ Simple YAML configuration
- ✅ Clean Code, no bloat, 100% Go

## 🛠 Usage

### 1. Install

```bash
git clone git@github.com:gBarczyszyn/mitosis.git
cd mitosis
go build -o mitosis
```

### 2. Create your `config.yaml`

```yaml
repo_url: git@github.com:gBarczyszyn/mitosis-dotfiles.git
tracked_paths:
  - ~/.zshrc
  - ~/.p10k.zsh
  - ~/.gitconfig
  - ~/.config/nvim/init.vim
```

### 3. Run sync

```bash
./mitosis sync --config config.yaml
```

### 4. Run apply

```bash
./mitosis apply --config config.yaml
```

### 5. Start the daemon

```bash
./mitosis daemon --config config.yaml
```

## 📦 Directory structure

By default, the repo is cloned into:

```
~/.mitosis/<repo-name>/
```

## 🧑‍💻 License

MIT © 2025 [Guilherme Barczyszyn](https://github.com/gBarczyszyn)
