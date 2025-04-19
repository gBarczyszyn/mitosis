#!/bin/bash

set -e

REPO="https://github.com/gBarczyszyn/mitosis.git"
INSTALL_DIR="/usr/local/bin"
BIN_NAME="mitosis"
SERVICE_NAME="mitosis.service"
SYSTEMD_USER_DIR="$HOME/.config/systemd/user"
CONFIG_PATH="$HOME/.mitosis/config.yaml"

echo "📥 Cloning mitosis..."
git clone $REPO /tmp/mitosis-install
cd /tmp/mitosis-install

echo "🔨 Building mitosis binary..."
go build -o $BIN_NAME

echo "🚀 Installing to $INSTALL_DIR..."
sudo mv $BIN_NAME $INSTALL_DIR/

echo "🛠 Setting up systemd user service..."

mkdir -p "$SYSTEMD_USER_DIR"

cat > "$SYSTEMD_USER_DIR/$SERVICE_NAME" <<EOF
[Unit]
Description=Mitosis Sync Daemon
After=network.target

[Service]
ExecStart=$INSTALL_DIR/mitosis daemon --config $CONFIG_PATH
Restart=always

[Install]
WantedBy=default.target
EOF

echo "🔄 Reloading systemd user services..."
systemctl --user daemon-reexec
systemctl --user enable $SERVICE_NAME
systemctl --user start $SERVICE_NAME

echo "✅ mitosis installed and running as a systemd user service!"
echo "👉 Place your config.yaml at: $CONFIG_PATH"
