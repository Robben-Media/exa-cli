# exa-cli

Exa AI Search CLI - Search, research, and get answers from the web using the [Exa](https://exa.ai) API.

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap builtbyrobben/tap
brew install exa-cli
```

### Download Binary

Download the latest release from [GitHub Releases](https://github.com/builtbyrobben/exa-cli/releases).

### Build from Source

```bash
git clone https://github.com/builtbyrobben/exa-cli.git
cd exa-cli
make build
```

## Authentication

### Set API Key

```bash
# Interactive (secure, recommended)
exa-cli auth set-key --stdin

# From environment variable
echo $EXA_API_KEY | exa-cli auth set-key --stdin

# From argument (discouraged - exposes in shell history)
exa-cli auth set-key YOUR_API_KEY
```

### Check Status

```bash
exa-cli auth status
```

### Remove Credentials

```bash
exa-cli auth remove
```

### Environment Variables

- `EXA_API_KEY` - Override stored credentials
- `EXA_CLI_KEYRING_BACKEND` - Force keyring backend (auto/keychain/file)
- `EXA_CLI_KEYRING_PASS` - Password for file backend (headless systems)

## Usage

### Search the web

```bash
exa-cli search "best practices for Go error handling"
exa-cli search "AI news" --num=10 --type=neural
exa-cli search "golang tutorials" --domains=go.dev,gobyexample.com
exa-cli search "climate research" --start-date=2024-01-01
```

### Search and retrieve page contents

```bash
exa-cli contents "how to use context in Go" --num=3
exa-cli contents "machine learning papers" --max-chars=5000 --highlights
exa-cli contents "API design patterns" --summary
```

### Find similar pages

```bash
exa-cli find-similar "https://go.dev/doc/effective_go" --num=5
exa-cli find-similar "https://example.com/article" --domains=dev.to,medium.com
```

### Get AI answers with citations

```bash
exa-cli answer "What is the difference between goroutines and threads?"
exa-cli answer "How does HTTP/3 improve performance?" --num=10
```

### JSON output (for scripting)

```bash
exa-cli search "Go concurrency" --json | jq '.results[].url'
exa-cli answer "What is WebAssembly?" --json
```

## Development

### Prerequisites

- Go 1.25+
- Make

### Commands

```bash
make build        # Build binary
make test         # Run tests
make lint         # Run linter
make fmt          # Format code
make ci           # Run full CI suite
make tools        # Install dev tools
```

## License

MIT

## Contributing

Contributions are welcome! Please read our contributing guidelines before submitting PRs.
