#!/bin/bash

set -e

BIN_NAME="mitosis"
INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.mitosis"
REPO_FILE="$CONFIG_DIR/repo.yaml"

echo "📦 Installing $BIN_NAME..."

# Move binary to /usr/local/bin
sudo cp "$BIN_NAME" "$INSTALL_DIR/"
sudo chmod +x "$INSTALL_DIR/$BIN_NAME"

# Create config dir
mkdir -p "$CONFIG_DIR"

# Only ask if file doesn't exist
if [ ! -f "$REPO_FILE" ]; then
  read -rp "🔗 Enter the Git repository URL to sync with: " REPO_URL

  echo "repo_url: $REPO_URL" > "$REPO_FILE"
  echo "✅ Repository URL saved to $REPO_FILE"
else
  echo "📁 Repo config already exists at $REPO_FILE, skipping..."
fi

read -rp "🚀 Do you want to start the mitosis daemon now? [y/N]: " START_DAEMON
if [[ "$START_DAEMON" =~ ^[Yy]$ ]]; then
    "$INSTALL_DIR/$BIN_NAME" start
fi
