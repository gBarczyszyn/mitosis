#!/bin/bash

set -e

echo "📦 Building mitosis..."
go build -o mitosis .

INSTALL_DIR="$HOME/.local/bin"
BIN_PATH="$INSTALL_DIR/mitosis"

mkdir -p "$INSTALL_DIR"
cp mitosis "$BIN_PATH"
chmod +x "$BIN_PATH"

echo "✅ Binary installed at $BIN_PATH"

# Ensure ~/.local/bin is in PATH
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "⚠️ $INSTALL_DIR is not in your PATH."

  SHELL_RC="${HOME}/.profile"
  [ -n "$ZSH_VERSION" ] && SHELL_RC="${HOME}/.zshrc"
  [ -n "$BASH_VERSION" ] && SHELL_RC="${HOME}/.bashrc"

  echo "export PATH=\"\$HOME/.local/bin:\$PATH\"" >> "$SHELL_RC"
  echo "➕ Added to $SHELL_RC"
  echo "🔁 Run 'source $SHELL_RC' or restart your terminal."
fi

CONFIG_DIR="$HOME/.mitosis"
mkdir -p "$CONFIG_DIR"

# Ask or use REPO_URL
if [ -z "$REPO_URL" ]; then
  read -rp "Enter the Git repository URL to sync with: " REPO_URL
fi

echo "repo_url: $REPO_URL" > "$CONFIG_DIR/repo.yaml"
echo "📁 repo.yaml saved at $CONFIG_DIR/repo.yaml"

# Create default config.yaml
echo "⚙️ Generating default config.yaml..."
"$BIN_PATH" init-config

echo "✅ Installation complete! You can now use 'mitosis'"
