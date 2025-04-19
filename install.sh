#!/bin/bash

set -e

REPO="https://github.com/gBarczyszyn/mitosis.git"
INSTALL_DIR="/usr/local/bin"
BIN_NAME="mitosis"
MITOSIS_DIR="$HOME/.mitosis"
CONFIG_PATH="$MITOSIS_DIR/config.yaml"

echo "📥 Cloning mitosis..."
git clone $REPO /tmp/mitosis-install
cd /tmp/mitosis-install

echo "🔨 Building mitosis binary..."
go build -o $BIN_NAME

echo "🚀 Installing to $INSTALL_DIR..."
sudo mv $BIN_NAME $INSTALL_DIR/

OS=$(uname -s)

if [[ "$OS" == "Darwin" ]]; then
  echo "🍎 Detected macOS - setting up launchctl service..."

  mkdir -p ~/Library/LaunchAgents

  cat > ~/Library/LaunchAgents/com.gbarczyszyn.mitosis.plist <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>com.gbarczyszyn.mitosis</string>
  <key>ProgramArguments</key>
  <array>
    <string>$INSTALL_DIR/mitosis</string>
    <string>daemon</string>
    <string>--config</string>
    <string>$CONFIG_PATH</string>
  </array>
  <key>RunAtLoad</key>
  <true/>
  <key>KeepAlive</key>
  <true/>
</dict>
</plist>
EOF

  launchctl unload ~/Library/LaunchAgents/com.gbarczyszyn.mitosis.plist 2>/dev/null || true
  launchctl load ~/Library/LaunchAgents/com.gbarczyszyn.mitosis.plist

  echo "✅ mitosis installed and running as a launchctl agent!"
elif [[ "$OS" == "Linux" ]] && command -v systemctl >/dev/null 2>&1; then
  echo "🐧 Detected Linux with systemd - setting up systemd user service..."

  SYSTEMD_USER_DIR="$HOME/.config/systemd/user"
  mkdir -p "$SYSTEMD_USER_DIR"

  cat > "$SYSTEMD_USER_DIR/mitosis.service" <<EOF
[Unit]
Description=Mitosis Sync Daemon
After=network.target

[Service]
ExecStart=$INSTALL_DIR/mitosis daemon --config $CONFIG_PATH
Restart=always

[Install]
WantedBy=default.target
EOF

  systemctl --user daemon-reexec
  systemctl --user enable mitosis.service
  systemctl --user start mitosis.service

  echo "✅ mitosis installed and running as a systemd user service!"
else
  echo "⚠️  Unsupported OS or no service manager found. Binary installed, but daemon not enabled."
fi

echo "👉 Place your config.yaml at: $CONFIG_PATH"
