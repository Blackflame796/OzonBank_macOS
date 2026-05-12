package ui

import (
	"OzonBank/internal/utils"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/webkit"
)

// WebView manages the application's web view
type WebView struct {
	logger  *utils.Logger
	webView webkit.WebView
	url     string
}

// NewWebView creates a new web view
func NewWebView(logger *utils.Logger, url string) *WebView {
	return &WebView{
		logger: logger,
		url:    url,
	}
}

// Create creates the web view component
func (wv *WebView) Create(frame foundation.Rect) appkit.View {
	config := webkit.NewWebViewConfiguration()
	webView := webkit.NewWebViewWithFrameConfiguration(frame, config)

	url := foundation.URL_URLWithString(wv.url)
	req := foundation.NewURLRequestWithURL(url)
	webView.LoadRequest(req)

	wv.webView = webView
	wv.logger.Info("WebView created and loaded URL: " + wv.url)

	return webView
}

// Reload reloads the web view
func (wv *WebView) Reload() {
	wv.webView.Reload(wv.webView)
	wv.logger.Info("WebView reloaded")
}

// LoadURL loads a specific URL
func (wv *WebView) LoadURL(urlString string) {
	url := foundation.URL_URLWithString(urlString)
	req := foundation.NewURLRequestWithURL(url)
	wv.webView.LoadRequest(req)
	wv.logger.Info("WebView loading URL: " + urlString)
}
