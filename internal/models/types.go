package models

// AppConfig represents application configuration
type AppConfig struct {
	Window WindowConfig
	Server ServerConfig
}

// WindowConfig contains window-specific settings
type WindowConfig struct {
	Width  int
	Height int
	Title  string
}

// ServerConfig contains server-specific settings
type ServerConfig struct {
	AuthURL string
}
