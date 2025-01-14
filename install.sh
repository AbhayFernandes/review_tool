#!/bin/bash
REPO="AbhayFernandes/review_tool"
VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep "tag_name" | cut -d '"' -f 4)
OS=$(uname -s)
ARCH=$(uname -m)

if [[ "$OS" == "Linux" ]]; then
  PLATFORM="linux"
elif [[ "$OS" == "Darwin" ]]; then
  PLATFORM="macos"
else
  echo "Unsupported OS"; exit 1
fi

URL="https://github.com/$REPO/releases/download/$VERSION/crev-$PLATFORM"
INSTALL_DIR="/usr/local/bin"

# Check for write permission
if [ ! -w "$INSTALL_DIR" ]; then
  echo "🔒 Permission denied for $INSTALL_DIR. Using sudo..."
  sudo curl -L "$URL" -o "$INSTALL_DIR/crev"
  sudo chmod +x "$INSTALL_DIR/crev"
else
  curl -L "$URL" -o "$INSTALL_DIR/crev"
  chmod +x "$INSTALL_DIR/crev"
fi

echo "✅ Installed crev version $VERSION"

