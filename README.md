# 🧬 mitosis

> Git-based workspace sync for UNIX devices, dotfiles-first.

**Mitosis** is a command-line tool to automatically sync your shell environment, dotfiles, and development setup across multiple Unix-based machines using Git.

## 🚀 Features

- ✅ Syncs system config files to a Git repo (e.g. `.zshrc`, `.gitconfig`, `~/.config/...`)
- ✅ Applies tracked files from Git into the system
- ✅ Watches for changes and auto-commits/pushes
- ✅ Simple YAML configuration stored in the repository
- ✅ Clean Code, no bloat, 100% Go
- ✅ Runs as a background systemd or launchctl service

## 🛠 Installation

### 🔹 Recommended one-liner:

```bash
REPO_URL=git@github.com:youruser/mitosis-gitops.git \
  bash <(curl -sL https://raw.githubusercontent.com/gBarczyszyn/mitosis/main/install.sh)
```

### 🔹 Interactive mode:

```bash
curl -sL https://raw.githubusercontent.com/gBarczyszyn/mitosis/main/install.sh -o install.sh
bash install.sh
```

This will:
- Clone the mitosis CLI
- Build and install the binary
- Run `mitosis init --repo <your-repo>`
- Start the daemon on macOS (`launchctl`) or Linux (`systemd`)

## 🧬 Usage

After installing, everything is driven by your Git-based `config.yaml` stored inside the repository itself:

```
~/.mitosis/<repo-name>/config.yaml
```

You can run:

```bash
mitosis sync        # Copy system files into repo
mitosis apply       # Apply repo files into system
mitosis daemon      # Run in watch mode
mitosis doctor      # Show resolved config
mitosis start       # Start the background service
mitosis stop        # Stop the background service
mitosis status      # Check service status
```

## 🛠 Development

To run locally:

```bash
make build
make sync
```

To initialize a new repository:

```bash
make init REPO=git@github.com:youruser/mitosis-gitops.git
```

## 📦 Directory structure

```
~/.mitosis/
└── mitosis-gitops/
    ├── config.yaml
    ├── .zshrc
    ├── .gitconfig
    ├── ...
```

## 🧑‍💻 License

MIT © 2025 [Guilherme Barczyszyn](https://github.com/gBarczyszyn)
