package ui

import (
	"OzonBank/internal/models"
	"OzonBank/internal/utils"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/objc"
)

// Window manages application window
type Window struct {
	config *models.AppConfig
	logger *utils.Logger
	window appkit.Window
}

// NewWindow creates a new window
func NewWindow(config *models.AppConfig, logger *utils.Logger) *Window {
	return &Window{
		config: config,
		logger: logger,
	}
}

// Create creates the application window
func (w *Window) Create(webView interface{}) {
	windowConfig := w.config.Window

	frame := appkit.NSRect{
		Size: appkit.NSSize{
			Width:  float64(windowConfig.Width),
			Height: float64(windowConfig.Height),
		},
	}

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		frame,
		appkit.ClosableWindowMask|appkit.TitledWindowMask,
		appkit.BackingStoreBuffered,
		false,
	)

	objc.Retain(&window)

	if wv, ok := webView.(appkit.View); ok {
		window.SetContentView(wv)
	}

	window.MakeKeyAndOrderFront(window)
	window.Center()
	window.SetTitle(windowConfig.Title)

	w.window = window
	w.logger.Info("Window created successfully")
}

// Show displays the window
func (w *Window) Show() {
	w.window.MakeKeyAndOrderFront(w.window)
	w.logger.Info("Window shown")
}

// Close closes the window
func (w *Window) Close() {
	w.window.Close()
	w.logger.Info("Window closed")
}
