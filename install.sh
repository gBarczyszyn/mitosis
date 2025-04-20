#!/bin/bash

set -e

REPO_URL="${REPO_URL:-}"

if [ -z "$REPO_URL" ]; then
  if [ -t 0 ]; then
    read -p "🌐 Enter your Git repository URL (e.g. git@github.com:user/mitosis-gitops.git): " REPO_URL
  else
    echo "❌ REPO_URL not provided and terminal not interactive. Please run manually or set REPO_URL env."
    exit 1
  fi
fi

REPO_NAME=$(basename "$REPO_URL" .git)
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="mitosis"

TMP_DIR=$(mktemp -d)
echo "📥 Cloning mitosis source into $TMP_DIR..."
git clone https://github.com/gBarczyszyn/mitosis.git "$TMP_DIR"
cd "$TMP_DIR"

echo "🔨 Building $BINARY_NAME..."
go build -o $BINARY_NAME

if [ ! -f "./$BINARY_NAME" ]; then
  echo "❌ Binary './$BINARY_NAME' not found."
  echo "👉 Build failed, check Go installation or source errors."
  exit 1
fi

echo "🚀 Installing to $INSTALL_DIR..."
sudo mv $BINARY_NAME $INSTALL_DIR/

# Ensure it's available in PATH
if ! command -v $BINARY_NAME >/dev/null 2>&1; then
  echo "⚠️  $BINARY_NAME not found in PATH. You may need to add $INSTALL_DIR to your PATH."
else
  echo "✅ $BINARY_NAME is now available at $(which $BINARY_NAME)"
fi

echo "📁 Running mitosis init..."
REPO_URL="$REPO_URL" $INSTALL_DIR/mitosis init --repo "$REPO_URL"

echo ""
read -p "🚀 Do you want to start the mitosis daemon now? (y/n): " RESP
if [[ "$RESP" == "y" || "$RESP" == "Y" ]]; then
  $INSTALL_DIR/mitosis start --config "$HOME/.mitosis/$REPO_NAME/config.yaml"
  echo "✅ Daemon started."
else
  echo "ℹ️  You can start the daemon later with: mitosis start"
fi
