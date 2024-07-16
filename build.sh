#!/bin/bash

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

OS=""
ARCH=""

case "$(uname -s)" in
    Darwin)
        OS="darwin"
        if [ "$(uname -m)" = "arm64" ]; then
            ARCH="arm64"
        else
            ARCH="amd64"
        fi
        ;;
    Linux)
        OS="linux"
        ARCH="amd64"
        ;;
    CYGWIN*|MINGW32*|MSYS*|MINGW*)
        OS="windows"
        ARCH="amd64"
        ;;
    *)
        echo -e "${RED}Unsupported OS${NC}"
        exit 1
        ;;
esac

# Set output binary name
OUTPUT="myterraform"

if [ "$OS" = "windows" ]; then
    OUTPUT="myterraform.exe"
fi

# Build the Go binary
echo -e "${YELLOW}Building for $OS/$ARCH...${NC}"
GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT main.go

if [ $? -ne 0 ]; then
    echo -e "${RED}Build failed${NC}"
    exit 1
fi

echo -e "${GREEN}Build successful: $OUTPUT${NC}"