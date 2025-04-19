#!/bin/bash

set -e

if [ -z "$REPO_URL" ]; then
  read -p "🌐 Enter your Git repository URL (e.g. git@github.com:user/mitosis-gitops.git): " REPO_URL
fi

REPO_NAME=$(basename "$REPO_URL" .git)
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="mitosis"

echo "📥 Cloning mitosis source..."
git clone https://github.com/gBarczyszyn/mitosis.git /tmp/mitosis-install
cd /tmp/mitosis-install

echo "🔨 Building $BINARY_NAME..."
go build -o $BINARY_NAME

echo "🚀 Installing to $INSTALL_DIR..."
sudo mv $BINARY_NAME $INSTALL_DIR/

echo "📁 Running mitosis init..."
$INSTALL_DIR/mitosis init --repo "$REPO_URL"

OS=$(uname -s)

if [[ "$OS" == "Darwin" ]]; then
  echo "🍎 Setting up launchctl daemon..."
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

elif [[ "$OS" == "Linux" ]] && command -v systemctl >/dev/null 2>&1; then
  echo "🐧 Setting up systemd user service..."
  SYSTEMD_USER_DIR="$HOME/.config/systemd/user"
  mkdir -p "$SYSTEMD_USER_DIR"

  CONFIG_PATH="$HOME/.mitosis/$REPO_NAME/config.yaml"

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

else
  echo "⚠️  Unsupported OS or no service manager found. Daemon mode not enabled."
fi

echo "✅ Mitosis installed and running!"
echo "📂 Repo path: ~/.mitosis/$REPO_NAME"
echo "📄 Config:    ~/.mitosis/$REPO_NAME/config.yaml"
