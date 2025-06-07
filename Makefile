VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.1.0")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

BINARY_NAME=olc
DIST_DIR=dist

# Build flags
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildTime=$(BUILD_TIME)"

.PHONY: all clean build build-all linux darwin windows

all: clean build-all

clean:
	rm -rf $(DIST_DIR)
	rm -f $(BINARY_NAME)

build:
	go build $(LDFLAGS) -o $(BINARY_NAME)

build-all: linux darwin windows

linux:
	@echo "Building for Linux..."
	@mkdir -p $(DIST_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64

darwin:
	@echo "Building for macOS..."
	@mkdir -p $(DIST_DIR)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64

windows:
	@echo "Building for Windows..."
	@mkdir -p $(DIST_DIR)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-arm64.exe

# Additional architectures
build-extended: build-all
	@echo "Building for additional architectures..."
	# FreeBSD
	GOOS=freebsd GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-freebsd-amd64
	# OpenBSD
	GOOS=openbsd GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-openbsd-amd64

# Create archives
archive: build-all
	@echo "Creating archives..."
	@cd $(DIST_DIR) && \
	for file in *; do \
		if [[ "$$file" == *".exe" ]]; then \
			zip "$${file%.exe}.zip" "$$file"; \
		else \
			tar -czf "$$file.tar.gz" "$$file"; \
		fi; \
	done

# Show available targets
list:
	@echo "Available targets:"
	@echo "  build       - Build for current platform"
	@echo "  build-all   - Build for Linux, macOS, Windows (amd64/arm64)"
	@echo "  linux       - Build for Linux (amd64/arm64)"
	@echo "  darwin      - Build for macOS (amd64/arm64)"
	@echo "  windows     - Build for Windows (amd64/arm64)"
	@echo "  build-extended - Build for additional platforms"
	@echo "  archive     - Create archives for distribution"
	@echo "  clean       - Remove build artifacts"
	@echo "  list        - Show this help"