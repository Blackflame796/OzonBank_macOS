# OzonBank Project Architecture

## Overview
This is a macOS desktop application built with Go that provides access to Ozon Finance services.

## Project Structure

```
OzonBank/
├── cmd/                          # Application entry points
│   └── ozonbank/
│       └── main.go              # Main entry point
│
├── internal/                     # Private application code
│   ├── app/                     # Core application logic
│   │   ├── app.go              # Main application struct and logic
│   │   └── config.go           # Configuration management
│   │
│   ├── ui/                      # User interface components
│   │   ├── window.go           # Window management
│   │   └── webview.go          # WebView management
│   │
│   ├── models/                  # Data structures and types
│   │   └── types.go            # Application models
│   │
│   └── utils/                   # Utility functions
│       └── logger.go            # Logging functionality
│
├── pkg/                          # Public reusable packages
│
├── config/                       # Configuration files
│
├── scripts/                      # Build and utility scripts
│   └── build.sh                 # Build script
│
├── docs/                         # Documentation
│
├── tests/                        # Integration tests
│
├── go.mod                        # Go modules definition
├── go.sum                        # Go modules checksums
└── README.md                     # Project documentation
```

## Architecture Principles

### 1. **Layered Architecture**
- **UI Layer** (internal/ui): Handles all macOS UI components (Window, WebView)
- **Application Layer** (internal/app): Contains business logic and application orchestration
- **Models Layer** (internal/models): Defines data structures
- **Utility Layer** (internal/utils): Cross-cutting concerns (logging)

### 2. **Dependency Injection**
Each component receives its dependencies through constructor functions:
- `NewApplication(config, logger)`
- `NewWindow(config, logger)`
- `NewWebView(logger, url)`

### 3. **Configuration Management**
- Centralized configuration in `internal/app/config.go`
- Easy to extend with external config files or environment variables

### 4. **Separation of Concerns**
- `cmd/`: Only handles application startup
- `internal/`: All business logic is internal to the application
- `pkg/`: Available for external packages (if needed in future)

## Building and Running

```bash
cd scripts
./build.sh
```

## Future Enhancements

1. **Add HTTP Server**: Create `internal/api` package for REST endpoints
2. **Add Data Persistence**: Create `internal/storage` package
3. **Add Tests**: Add unit tests in `tests/` directory
4. **Add Configuration Files**: Support YAML/JSON config files in `config/`
5. **Add Error Handling**: Enhance error handling across all packages
6. **Add Middleware**: Add middleware for logging, error handling, etc.
