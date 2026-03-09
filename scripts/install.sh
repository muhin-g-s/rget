#!/bin/bash

APP_NAME="rget"
BIN_PATH="$HOME/bin"
OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

mkdir -p "$BIN_PATH"

LATEST=$(curl -s https://api.github.com/repos/muhin-g-s/rget/releases/latest \
         | grep '"tag_name"' | head -n1 | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')

if [ -z "$LATEST" ]; then
    echo "Unable to get latest version of  $APP_NAME"
    exit 1
fi

BINARY_URL="https://github.com/muhin-g-s/rget/releases/download/$LATEST/${APP_NAME}-${OS}-${ARCH}"

curl -L "$BINARY_URL" -o "$BIN_PATH/$APP_NAME"
chmod +x "$BIN_PATH/$APP_NAME"

if [[ ":$PATH:" != *":$BIN_PATH:"* ]]; then
    export PATH="$BIN_PATH:$PATH"
    SHELL_RC="$HOME/.zshrc"
    [ -f "$HOME/.bashrc" ] && SHELL_RC="$HOME/.bashrc"
    echo "export PATH=\"$BIN_PATH:\$PATH\"" >> "$SHELL_RC"
fi

echo "$APP_NAME installed in $BIN_PATH version $LATEST"
