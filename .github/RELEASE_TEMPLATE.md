# Release Template

Use this template when manually creating releases.

## Release Notes for OLC v{VERSION}

### What's New
- Feature 1
- Feature 2
- Bug fix 1

### Downloads

Choose the appropriate binary for your platform:

| Platform | Architecture | Download |
|----------|--------------|----------|
| Linux | x86_64 | [olc-linux-amd64.tar.gz](download-link) |
| Linux | ARM64 | [olc-linux-arm64.tar.gz](download-link) |
| macOS | Intel | [olc-darwin-amd64.tar.gz](download-link) |
| macOS | Apple Silicon | [olc-darwin-arm64.tar.gz](download-link) |
| Windows | x86_64 | [olc-windows-amd64.zip](download-link) |
| Windows | ARM64 | [olc-windows-arm64.zip](download-link) |

### Installation

1. Download the appropriate archive for your platform
2. Extract the binary:
   ```bash
   # Linux/macOS
   tar -xzf olc-*.tar.gz
   
   # Windows
   unzip olc-*.zip
   ```
3. Move to a directory in your PATH:
   ```bash
   # Linux/macOS
   sudo mv olc /usr/local/bin/
   
   # Or for user-only installation
   mv olc ~/.local/bin/
   ```

### Quick Start

```bash
# Configure your Ollama server
olc config set ip your-ollama-server-ip
olc config set model your-preferred-model

# Start chatting
olc chat

# Generate text
olc generate --prompt "Your prompt here"

# Manage models
olc model list
```

### Verification

Verify the integrity of your download using SHA256 checksums:
```bash
sha256sum olc-*
```

Expected checksums:
```
{CHECKSUMS_WILL_BE_HERE}
```