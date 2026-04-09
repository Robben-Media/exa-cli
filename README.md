# exa-cli

Command-line interface for the [Exa](https://exa.ai) AI search API. Search the web, retrieve page contents, find similar pages, and get AI-powered answers with citations.

## Installation

### Download Binary

Download the latest release from [GitHub Releases](https://github.com/builtbyrobben/exa-cli/releases).

### Build from Source

```bash
git clone https://github.com/builtbyrobben/exa-cli.git
cd exa-cli
make build
```

## Configuration

exa-cli authenticates via an Exa API key. You can provide it in two ways:

**Environment variable (recommended for CI/scripts):**

```bash
export EXA_API_KEY="your-api-key"
```

**Keyring storage (recommended for interactive use):**

```bash
# Interactive prompt (secure)
exa-cli auth set-key --stdin

# Pipe from environment
echo "$EXA_API_KEY" | exa-cli auth set-key --stdin
```

### Environment Variables

| Variable | Description |
|----------|-------------|
| `EXA_API_KEY` | API key (overrides keyring) |
| `EXA_CLI_COLOR` | Color output: `auto`, `always`, `never` |
| `EXA_CLI_OUTPUT` | Default output mode: `json`, `plain` |

## Global Flags

| Flag | Description |
|------|-------------|
| `--json` | Output JSON to stdout (best for scripting) |
| `--plain` | Output stable, parseable text (TSV; no colors) |
| `--color` | Color output: `auto`, `always`, `never` |
| `--verbose` | Enable verbose logging |
| `--force` | Skip confirmations for destructive commands |
| `--no-input` | Never prompt; fail instead (useful for CI) |

## Commands

### auth

Manage authentication credentials.

```bash
# Store API key in system keyring
exa-cli auth set-key --stdin

# Check authentication status
exa-cli auth status

# Remove stored credentials
exa-cli auth remove
```

### search

Search the web using Exa's neural or keyword search.

```bash
# Basic search
exa-cli search "best practices for Go error handling"

# Control number of results
exa-cli search "AI news" --num 10

# Search type: auto, neural, or keyword
exa-cli search "golang tutorials" --type neural

# Filter by domain
exa-cli search "golang tutorials" --domains go.dev,gobyexample.com

# Exclude domains
exa-cli search "tech news" --exclude-domains reddit.com,twitter.com

# Filter by publish date
exa-cli search "climate research" --start-date 2024-01-01 --end-date 2024-12-31

# Output as JSON
exa-cli search "Go concurrency" --json
```

### contents

Search and retrieve full page contents in a single call.

```bash
# Search with page text
exa-cli contents "how to use context in Go"

# Control number of results and text length
exa-cli contents "machine learning papers" --num 5 --max-chars 5000

# Include highlights
exa-cli contents "API design patterns" --highlights

# Include AI summary per result
exa-cli contents "distributed systems" --summary

# Filter by domain
exa-cli contents "React hooks" --domains react.dev,dev.to

# Output as JSON
exa-cli contents "WebAssembly tutorial" --json
```

### find-similar

Find pages similar to a given URL.

```bash
# Find similar pages
exa-cli find-similar "https://go.dev/doc/effective_go"

# Control number of results
exa-cli find-similar "https://example.com/article" --num 10

# Filter by domain
exa-cli find-similar "https://example.com" --domains dev.to,medium.com

# Exclude domains
exa-cli find-similar "https://example.com" --exclude-domains reddit.com

# Output as JSON
exa-cli find-similar "https://go.dev/blog" --json
```

### answer

Get AI-powered answers with source citations.

```bash
# Ask a question
exa-cli answer "What is the difference between goroutines and threads?"

# Control number of source results
exa-cli answer "How does HTTP/3 improve performance?" --num 10

# Output as JSON
exa-cli answer "What is WebAssembly?" --json
```

### version

Print version information.

```bash
exa-cli version
```

## License

MIT
