# Project Architecture Diagram

## Overall Architecture

```
┌────────────────────────────────────────────────────────────────┐
│                    OzonBank Application                        │
└────────────────────────────────────────────────────────────────┘

┌──────────────────────── Entry Point ────────────────────────────┐
│  cmd/ozonbank/main.go                                          │
│  • Initialize Logger                                            │
│  • Load Configuration                                           │
│  • Create and Run Application                                   │
└────────────────────────────────────────────────────────────────┘
                               ↓
┌──────────────────── Application Layer ──────────────────────────┐
│  internal/app/app.go                                            │
│  • Application struct                                           │
│  • Run method for starting the app                              │
│  • Orchestration of UI and WebView components                   │
└────────────────────────────────────────────────────────────────┘
                      ↙              ↘
┌──────────────────────────────────────────────────────────────────┐
│  UI Layer                          │          Configuration Layer │
│  ────────────────────────────────  │          ──────────────────── │
│                                    │                               │
│  • internal/ui/window.go           │  internal/app/config.go      │
│    - Window struct                 │  • LoadConfig()              │
│    - Create/Show/Close methods     │  • Returns AppConfig         │
│                                    │                               │
│  • internal/ui/webview.go          │                               │
│    - WebView struct                │                               │
│    - Create/Reload/LoadURL methods │                               │
└────────────────┬──────────────────────────────────────┬──────────┘
                 │                                      │
                 ↓                                      ↓
┌─────────────────────────────────────────────────────────────────┐
│                      Models & Data Layer                        │
│  ──────────────────────────────────────────────────────────────  │
│                                                                  │
│  internal/models/types.go                                       │
│  • AppConfig struct                                             │
│  • WindowConfig struct                                          │
│  • ServerConfig struct                                          │
└──────────────┬──────────────────────────────────────────────────┘
               │
               ↓
┌─────────────────────────────────────────────────────────────────┐
│                    Utilities & External APIs                    │
│  ──────────────────────────────────────────────────────────────  │
│                                                                  │
│  internal/utils/logger.go                                       │
│  • Logger struct                                                │
│  • Info/Warn/Error/Debug methods                                │
│                                                                  │
│  External:                                                      │
│  • github.com/progrium/darwinkit - macOS APIs                   │
│  • github.com/webview/webview_go - WebView                      │
└─────────────────────────────────────────────────────────────────┘
```

## Package Dependencies

```
cmd/ozonbank/main.go
  ├── internal/app
  ├── internal/utils (Logger)
  └── [depends on]

internal/app/app.go
  ├── internal/models (AppConfig)
  ├── internal/ui (Window, WebView)
  ├── internal/utils (Logger)
  └── External: darwinkit, foundation, appkit

internal/app/config.go
  └── internal/models (AppConfig)

internal/ui/window.go
  ├── internal/models (AppConfig)
  ├── internal/utils (Logger)
  └── External: darwinkit, appkit

internal/ui/webview.go
  ├── internal/utils (Logger)
  └── External: appkit, foundation, webkit

internal/models/types.go
  └── [no dependencies - base data types]

internal/utils/logger.go
  └── [no dependencies - only stdlib]
```

## Data Flow

```
1. Application Startup
   └─► Initialize Logger
       └─► Load Configuration
           └─► Create Application Instance
               └─► Run Application
                   ├─► macOS RunApp event loop
                   ├─► Create WebView with auth URL
                   ├─► Create Window with WebView
                   └─► Start event loop

2. User Interaction
   └─► macOS Event Loop
       ├─► Window events (close, resize, etc.)
       └─► WebView events (navigation, content loading)

3. Application Termination
   └─► Window closed
       └─► Application delegate triggers
           └─► Application terminates
```

## File Organization

```
OzonBank/
│
├── cmd/                              # Application entry points
│   └── ozonbank/
│       └── main.go                  # ┐
│                                    │ Minimal startup code
├── internal/                        │ (15 lines)
│   ├── app/                         │
│   │   ├── app.go                   ├─ Core logic
│   │   └── config.go                │ (Orchestration)
│   │                                │
│   ├── ui/                          ├─ UI components
│   │   ├── window.go                │ (View layer)
│   │   └── webview.go               │
│   │                                │
│   ├── models/                      ├─ Data structures
│   │   └── types.go                 │ (Immutable definitions)
│   │                                │
│   └── utils/                       └─ Cross-cutting concerns
│       └── logger.go                  (Logging)
│
├── config/
│   └── config.yaml                  # Configuration template
│
├── scripts/
│   └── build.sh                     # Build automation
│
├── docs/
│   ├── DEVELOPMENT.md               # Development guide
│   ├── RESOURCES.md                 # Reference materials
│   └── ARCHITECTURE.md              # This file
│
├── tests/
│   └── integration_test.go          # Test cases
│
├── go.mod                           # Module definition
├── go.sum                           # Dependency checksums
├── README.md                        # Project overview
├── CONTRIBUTING.md                  # Contributing guide
└── .gitignore                       # Git ignore rules
```

## Key Principles

### 1. Layered Architecture
- **Entry Point**: Minimal, only initialization
- **Application**: Orchestration and coordination
- **UI**: All visual components
- **Models**: Pure data structures
- **Utils**: Reusable functionality

### 2. Dependency Injection
All dependencies passed through constructors:
```go
NewApplication(config, logger)
NewWindow(config, logger)
NewWebView(logger, url)
```

### 3. Single Responsibility
- Each package has one clear purpose
- No circular dependencies
- Clear separation of concerns

### 4. Extensibility
- Easy to add new packages (e.g., `internal/storage`, `internal/api`)
- Configuration is centralized and easy to modify
- Logger can be replaced with more advanced implementations

## Scale Potential

This architecture can easily grow to handle:
- Multiple windows and views
- Background services and workers
- Database persistence layer
- REST API server
- WebSocket connections
- Advanced error handling
- Comprehensive testing
