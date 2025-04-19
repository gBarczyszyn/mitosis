#!/bin/bash

set -e

BIN_NAME="mitosis"
BIN_PATH="./$BIN_NAME"
INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.mitosis"
REPO_FILE="$CONFIG_DIR/repo.yaml"

echo "📦 Installing $BIN_NAME..."

# Validar se o binário existe
if [ ! -f "$BIN_PATH" ]; then
  echo "❌ Binary '$BIN_PATH' not found."
  echo "👉 Run 'go build -o mitosis .' before running this script."
  exit 1
fi

# Instalar binário
sudo cp "$BIN_PATH" "$INSTALL_DIR/"
sudo chmod +x "$INSTALL_DIR/$BIN_NAME"
echo "✅ Installed $BIN_NAME to $INSTALL_DIR"

# Criar pasta de config
mkdir -p "$CONFIG_DIR"

# Se o repo.yaml não existir, criar
if [ ! -f "$REPO_FILE" ]; then
  if [ -n "$REPO_URL" ]; then
    echo "repo_url: $REPO_URL" > "$REPO_FILE"
  else
    read -rp "🔗 Enter the Git repository URL to sync with: " REPO_URL_INPUT
    echo "repo_url: $REPO_URL_INPUT" > "$REPO_FILE"
  fi
  echo "✅ Repository URL saved to $REPO_FILE"
else
  echo "📁 Repo config already exists at $REPO_FILE, skipping..."
fi

# Perguntar se deseja iniciar o daemon
read -rp "🚀 Do you want to start the mitosis daemon now? [y/N]: " START_DAEMON
if [[ "$START_DAEMON" =~ ^[Yy]$ ]]; then
    "$INSTALL_DIR/$BIN_NAME" start
fi
