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
curl -L "$URL" -o /usr/local/bin/crev
chmod +x /usr/local/bin/crev
echo "Installed crev version $VERSION"

