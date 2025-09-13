# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project Overview

This is a high-performance network proxy system written in Go that provides
multi-protocol support including SOCKS5, WebSocket, and HTTP proxy
functionality. The system handles protocol conversion, dynamic upstream
connection selection, authentication, and configuration hot reload.

## Essential Commands

```bash
# Build the project
go build -o proxy-server cmd/main.go

# Run the proxy server
go run cmd/main.go

# Run with custom configuration
go run cmd/main.go -config config/server-config.json

# Run tests
go test ./... -v

# Format code and run tests
go fmt ./... && go test ./... && go mod tidy

# Development workflow
go mod tidy && go fmt ./... && go test ./... -v
```

## Architecture Overview

### Core Structure

- **`cmd/main.go`** - Main entry point with command line parsing and server
  initialization
- **`pkg/interfaces/`** - Unified interface definitions for clients and servers
- **Protocol Packages** (`pkg/socks5/`, `pkg/websocket/`, `pkg/http/`) -
  Protocol-specific implementations
- **`pkg/upstream/`** - Dynamic upstream connection selection with multiple
  strategies
- **`pkg/config/`** - Configuration management with file watching and hot reload
- **`pkg/proxy/`** - Factory pattern for protocol-specific instance creation

### Key Patterns

- **Factory Pattern**: Protocol-specific client/server creation in
  `pkg/proxy/factory.go`
- **Strategy Pattern**: Upstream connection selection (round-robin, random,
  weighted, failover)
- **Observer Pattern**: Configuration hot reload with event notifications
- **Interface Segregation**: Clean separation between client and server
  interfaces

## Configuration

The system supports flexible configuration through:

- Command line parameters for quick setup
- JSON configuration files for complex scenarios
- Hot reload support for runtime changes

Example configuration includes upstream servers, authentication methods, logging
levels, and protocol-specific settings.

## Protocol Support

### SOCKS5 Implementation

- Full RFC 1928/1929 compliance
- Username/password authentication
- IPv4 and domain address support
- Connection handling in `pkg/socks5/server.go:95`

### WebSocket Support

- Custom header authentication
- Protocol conversion from SOCKS5
- Real-time data forwarding

### HTTP Proxy

- CONNECT method support
- Address format: `http://proxy-server:port`

## Address Format Support

The system supports various address formats:

- `tcp://host:port` - Standard TCP
- `tls://host:port` - TCP with TLS
- `ws://host:port` - WebSocket
- `wss://host:port` - WebSocket with TLS
- `socks5://host:port` - SOCKS5 proxy
- `http://host:port` - HTTP proxy

## Development Guidelines

### Code Standards

- Use `go fmt` for code formatting
- Program to interfaces, not implementations
- Implement comprehensive error handling
- Write tests for new features and bug fixes

### Testing Strategy

- Unit tests in individual packages (`go test ./pkg/socks5`)
- Integration tests in `tests/` directory
- Test real network connections and concurrent access
- Use test coverage analysis: `go test -cover ./...`

### Common Development Tasks

**Adding New Protocol Support:**

1. Create new package under `pkg/`
2. Implement `ClientInterface` and `ServerInterface` from `pkg/interfaces/`
3. Register in `pkg/proxy/factory.go`
4. Add address format parsing

**Adding Upstream Strategy:**

1. Implement `UpstreamSelector` interface in `pkg/upstream/`
2. Add strategy type to `pkg/upstream/strategies.go`
3. Update configuration schema

**Configuration Changes:**

1. Modify struct definitions in `pkg/config/config.go`
2. Update validation in `pkg/config/validate.go`
3. Add examples to configuration files

## Important Files

- `cmd/main.go:main()` - Application entry point
- `pkg/proxy/factory.go:CreateServer()` - Protocol server creation
- `pkg/upstream/manager.go:GetConnection()` - Upstream selection logic
- `pkg/config/config.go:LoadConfiguration()` - Configuration loading
- `pkg/interfaces/` - Core interface definitions
