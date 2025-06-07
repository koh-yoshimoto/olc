#!/bin/bash

# Build script for cross-platform compilation
set -e

VERSION=${VERSION:-$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.1.0")}
COMMIT=${COMMIT:-$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")}
BUILD_TIME=${BUILD_TIME:-$(date -u +"%Y-%m-%dT%H:%M:%SZ")}

BINARY_NAME="olc"
DIST_DIR="dist"

# Build flags
LDFLAGS="-s -w -X main.version=$VERSION -X main.commit=$COMMIT -X main.buildTime=$BUILD_TIME"

echo "Building OLC (Ollama Client) v$VERSION ($COMMIT)"
echo "Build time: $BUILD_TIME"
echo

# Clean previous builds
rm -rf $DIST_DIR
mkdir -p $DIST_DIR

# Build matrix
declare -A platforms=(
    ["linux/amd64"]="$DIST_DIR/$BINARY_NAME-linux-amd64"
    ["linux/arm64"]="$DIST_DIR/$BINARY_NAME-linux-arm64"
    ["darwin/amd64"]="$DIST_DIR/$BINARY_NAME-darwin-amd64"
    ["darwin/arm64"]="$DIST_DIR/$BINARY_NAME-darwin-arm64"
    ["windows/amd64"]="$DIST_DIR/$BINARY_NAME-windows-amd64.exe"
    ["windows/arm64"]="$DIST_DIR/$BINARY_NAME-windows-arm64.exe"
)

# Build for each platform
for platform in "${!platforms[@]}"; do
    IFS='/' read -r GOOS GOARCH <<< "$platform"
    output="${platforms[$platform]}"
    
    echo "Building for $GOOS/$GOARCH..."
    
    if ! GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "$LDFLAGS" -o "$output"; then
        echo "Error: Failed to build for $GOOS/$GOARCH"
        exit 1
    fi
    
    # Calculate file size
    if [[ "$GOOS" == "windows" ]]; then
        size=$(stat -f%z "$output" 2>/dev/null || stat -c%s "$output" 2>/dev/null || echo "unknown")
    else
        size=$(stat -f%z "$output" 2>/dev/null || stat -c%s "$output" 2>/dev/null || echo "unknown")
    fi
    
    # Convert bytes to human readable
    if [[ "$size" != "unknown" ]]; then
        if (( size > 1048576 )); then
            size_mb=$(awk "BEGIN {printf \"%.1f\", $size/1048576}")
            size_str="${size_mb}MB"
        elif (( size > 1024 )); then
            size_kb=$(awk "BEGIN {printf \"%.1f\", $size/1024}")
            size_str="${size_kb}KB"
        else
            size_str="${size}B"
        fi
    else
        size_str="unknown"
    fi
    
    echo "✓ Built $output ($size_str)"
done

echo
echo "Build complete! Binaries are in the $DIST_DIR directory:"
ls -la $DIST_DIR/

echo
echo "To create archives for distribution, run:"
echo "  make archive"