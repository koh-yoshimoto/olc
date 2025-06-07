# OLC (Ollama Client)

OLC is a sleek command-line interface client for interacting with Ollama API in private network environments.

## Installation

### Pre-built Binaries

Download the latest release for your platform from the [releases page](https://github.com/koh/ollama-client/releases).

| Platform | Architecture | Download |
|----------|--------------|----------|
| Linux | x86_64 | Download from releases |
| Linux | ARM64 | Download from releases |
| macOS | Intel | Download from releases |
| macOS | Apple Silicon | Download from releases |
| Windows | x86_64 | Download from releases |
| Windows | ARM64 | Download from releases |

**Installation:**
1. Download the appropriate archive
2. Extract: `tar -xzf olc-*.tar.gz` (Linux/macOS) or `unzip olc-*.zip` (Windows)
3. Move to PATH: `sudo mv olc /usr/local/bin/` or place in your preferred directory

### Build from Source

```bash
git clone https://github.com/koh/ollama-client.git
cd ollama-client

# Build for current platform
go build -o olc

# Or build for all platforms
make build-all

# Create distribution archives
make archive
```

### Available Platforms

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64, arm64)

Build commands:
```bash
make build        # Current platform
make build-all     # All platforms
make linux         # Linux only
make darwin        # macOS only  
make windows       # Windows only
make archive       # Create archives
```

## Configuration

### Configuration Commands

Set IP address:
```bash
olc config set ip 192.168.1.100
```

Set port (default: 11434):
```bash
olc config set port 11434
```

Set default model:
```bash
olc config set model mistral-small:latest
```

Show current configuration:
```bash
olc config show
```

### Configuration File

Configuration is automatically saved to `~/.olc.yaml`:

```yaml
ip: 192.168.1.100
port: "11434"
default_model: mistral-small:latest
```

Default values:
- IP: `localhost`
- Port: `11434`

## Usage

### Global Options

- `--config`: Path to config file

### Chat Command

Start an interactive chat session:

```bash
# Basic chat (uses default model if set)
olc chat

# With specific model
olc chat --model llama2

# With system prompt
olc chat --system "You are a helpful assistant"

# With custom temperature
olc chat --temperature 0.9
```

### Generate Command

Generate text completion:

```bash
# Basic generation (uses default model if set)
olc generate --prompt "Write a haiku about programming"

# With specific model and temperature
olc generate --model llama2 --prompt "Explain quantum computing" --temperature 0.5
```

### Model Management

List available models:

```bash
olc model list
# or
olc model ls
```

Set default model (deprecated - use config set model):

```bash
olc model set mistral-small:latest
```

Pull a new model:

```bash
olc model pull llama2:7b
```

Delete a model:

```bash
olc model delete llama2:7b
# or
olc model rm llama2:7b
```

## Examples

### Using with Private Network

```bash
# Configure IP address for private network
olc config set ip 192.168.1.100

# Use default port (11434) or set custom port
olc config set port 8080

# Start chatting
olc chat
```

### Batch Processing

```bash
# Generate response and save to file
olc generate --prompt "$(cat prompt.txt)" > response.txt

# Process multiple prompts
for prompt in prompts/*.txt; do
    olc generate --prompt "$(cat $prompt)" > "responses/$(basename $prompt)"
done
```

## Features

- **Interactive Chat**: Multi-turn conversations with context retention
- **Text Generation**: Single-shot text completion
- **Model Management**: List, pull, and delete models
- **Performance Metrics**: Display token generation speed
- **Flexible Configuration**: Support for config files, environment variables, and command-line flags
- **Private Network Support**: Easily configure custom Ollama API endpoints

## Requirements

- Go 1.21 or later
- Access to an Ollama API endpoint

## License

MIT License