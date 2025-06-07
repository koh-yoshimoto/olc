# Contributing to OLC

Thank you for your interest in contributing to OLC (Ollama Client)!

## Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/koh/ollama-client.git
   cd ollama-client
   ```

2. **Install Go 1.21+**
   Make sure you have Go 1.21 or later installed.

3. **Build and test**
   ```bash
   # Build for current platform
   go build -o olc
   
   # Run tests
   go test -v ./...
   
   # Build for all platforms
   make build-all
   ```

## Making Changes

1. **Create a branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Follow Go best practices
   - Add tests for new functionality
   - Update documentation as needed

3. **Test your changes**
   ```bash
   go test -v ./...
   go build -o olc
   ./olc --help
   ```

4. **Create a pull request**
   - Describe your changes clearly
   - Reference any related issues

## Release Process

Releases are automated through GitHub Actions:

1. **Create a tag**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **GitHub Actions will:**
   - Build binaries for all platforms
   - Create archives
   - Generate checksums
   - Create a GitHub release with all assets

## Manual Release (if needed)

If you need to create a release manually:

1. **Build all platforms**
   ```bash
   make build-all
   make archive
   ```

2. **Create checksums**
   ```bash
   cd dist
   sha256sum * > checksums.txt
   ```

3. **Create GitHub release**
   - Go to GitHub releases page
   - Create new release
   - Upload all files from `dist/` directory
   - Use the template from `.github/RELEASE_TEMPLATE.md`

## Code Style

- Follow `gofmt` formatting
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions focused and small

## Testing

- Write tests for new functionality
- Ensure all tests pass before submitting PR
- Test on multiple platforms when possible