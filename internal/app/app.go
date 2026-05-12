package app

import (
	"OzonBank/internal/models"
	"OzonBank/internal/ui"
	"OzonBank/internal/utils"

	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

// Application represents the main application
type Application struct {
	config  *models.AppConfig
	logger  *utils.Logger
	window  *ui.Window
	webView *ui.WebView
}

// NewApplication creates a new application instance
func NewApplication(config *models.AppConfig, logger *utils.Logger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}

// Run starts the application
func (a *Application) Run() {
	a.logger.Info("Starting OzonBank application")

	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		a.logger.Info("macOS application initialized")

		// Set application properties
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create WebView
		frame := foundation.Rect{
			Size: foundation.Size{
				Width:  float64(a.config.Window.Width),
				Height: float64(a.config.Window.Height),
			},
		}

		a.webView = ui.NewWebView(a.logger, a.config.Server.AuthURL)
		webViewComponent := a.webView.Create(frame)

		// Create Window
		a.window = ui.NewWindow(a.config, a.logger)
		a.window.Create(webViewComponent)

		// Set delegate behavior
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})

		a.logger.Info("OzonBank application started successfully")
	})
}
