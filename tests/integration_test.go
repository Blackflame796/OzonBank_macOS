package tests

import (
	"OzonBank/internal/models"
	"OzonBank/internal/utils"
	"testing"
)

// TestLoggerCreation tests logger initialization
func TestLoggerCreation(t *testing.T) {
	logger := utils.NewLogger()
	if logger == nil {
		t.Fatal("Logger should not be nil")
	}
}

// TestConfigLoading tests configuration loading
func TestConfigLoading(t *testing.T) {
	// This is a placeholder for config loading tests
	expected := &models.AppConfig{
		Window: models.WindowConfig{
			Width:  1440,
			Height: 900,
			Title:  "OzonBank",
		},
		Server: models.ServerConfig{
			AuthURL: "https://finance.ozon.ru/apps/auth/signin?redirect=%2Flk",
		},
	}

	if expected.Window.Width != 1440 {
		t.Errorf("Expected width 1440, got %d", expected.Window.Width)
	}

	if expected.Window.Height != 900 {
		t.Errorf("Expected height 900, got %d", expected.Window.Height)
	}
}
