package app

import (
	"OzonBank/internal/models"
)

// LoadConfig loads application configuration
func LoadConfig() *models.AppConfig {
	return &models.AppConfig{
		Window: models.WindowConfig{
			Width:  1440,
			Height: 900,
			Title:  "OzonBank",
		},
		Server: models.ServerConfig{
			AuthURL: "https://finance.ozon.ru/apps/auth/signin?redirect=%2Flk",
		},
	}
}
