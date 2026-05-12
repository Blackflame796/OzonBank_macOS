# Development Plan

## Architecture Overview

The OzonBank project follows a layered architecture pattern with clear separation of concerns:

```
┌─────────────────────────────────────┐
│      Entry Point (cmd/ozonbank)     │
├─────────────────────────────────────┤
│     Application Layer (app)         │
│  - Orchestration                    │
│  - Configuration                    │
├─────────────────────────────────────┤
│  UI Layer (ui) & Models (models)    │
│  - Window management                │
│  - WebView management               │
│  - Data structures                  │
├─────────────────────────────────────┤
│     Utilities & External APIs       │
│  - Logging                          │
│  - DarwinKit (macOS)                │
│  - WebKit                           │
└─────────────────────────────────────┘
```

## Key Design Patterns Used

### 1. Dependency Injection
All dependencies are passed through constructors:
```go
func NewApplication(config *AppConfig, logger *Logger) *Application
```

### 2. Single Responsibility Principle
Each package has a specific responsibility:
- `app`: Application orchestration
- `ui`: UI components
- `models`: Data definitions
- `utils`: Cross-cutting concerns

### 3. Logging
Centralized logging through `utils.Logger`:
- Info: General information
- Warn: Warning messages
- Error: Error messages
- Debug: Debug information

## Future Improvements

### Phase 1: Enhanced Configuration
- [ ] Support YAML/JSON config files
- [ ] Environment variable support
- [ ] Configuration validation

### Phase 2: Error Handling
- [ ] Custom error types
- [ ] Error recovery mechanisms
- [ ] Better error messages

### Phase 3: Features
- [ ] Add API server for inter-process communication
- [ ] Add data persistence layer
- [ ] Add user authentication
- [ ] Add data caching

### Phase 4: Testing
- [ ] Unit tests for all packages
- [ ] Integration tests
- [ ] Mock objects for external dependencies

### Phase 5: Monitoring
- [ ] Enhanced logging with different levels
- [ ] Performance monitoring
- [ ] Error tracking

## Dependencies

Current dependencies in `go.mod`:
- `github.com/progrium/darwinkit`: macOS application framework
- `github.com/webview/webview_go`: WebView support

## Build Instructions

```bash
# Navigate to scripts directory
cd scripts

# Run build script
./build.sh

# The binary will be in the root directory
./main
```

## Coding Standards

1. **Naming**: Use PascalCase for exported identifiers, camelCase for private
2. **Comments**: Add godoc comments for all exported types and functions
3. **Error Handling**: Handle errors explicitly, don't ignore them
4. **Logging**: Use logger for all important events and errors
5. **Testing**: Write tests for new functionality
