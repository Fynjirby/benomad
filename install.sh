#!/bin/bash

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
FILE="unknown"

case "$OS-$ARCH" in
    linux-x86_64)   FILE="benomad-linux-amd" ;;
    linux-aarch64)  FILE="benomad-linux-arm" ;;
    darwin-arm64)   FILE="benomad-macos-arm" ;;
    darwin-x86_64)   FILE="benomad-macos-amd" ;;
    windows-x86_64)   FILE="benomad-windows-amd" ;;
    windows-aarch64)   FILE="benomad-windows-arm" ;;
    *)
        echo "Unsupported OS/arch: $OS-$ARCH"
        exit 1
        ;;
esac

URL="https://github.com/Fynjirby/benomad/releases/download/benomad/$FILE"
curl -L -O "$URL" || { echo "Download failed"; exit 1; }

chmod +x "$FILE"
echo "Successfully downloaded $FILE!"
echo "Moving to /usr/local/bin, may require password"
echo "If this fails please move the file to some bin directory in your PATH yourself"
sudo mv $FILE /usr/local/bin/benomad
echo "Successfully moved to /usr/local/bin, now run with a simple *benomad*"
