package common

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"time"
)

// ServerConfig configures gin server.
type ServerConfig struct {
	Host string
	Port string

	GinMode string

	Limits     []string
	Operations map[string]string
}

// DatabaseConfig stores DB credentials.
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

const (
	defaultHost     = "localhost:8080"
	defaultBasePath = "/api/v1"
)

var defaultSchemes = []string{"http"}

// SwaggerConfig configures swaggo/swag.
type SwaggerConfig struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
}

// NewSwaggerConfig returns *SwaggerConfig with preconfigured fields.
func NewSwaggerConfig(title, description, version string) *SwaggerConfig {
	return &SwaggerConfig{
		Title:       title,
		Description: description,
		Version:     version,
		Host:        defaultHost,
		BasePath:    defaultBasePath,
		Schemes:     defaultSchemes,
	}
}

// DefaultCorsConfig returns cors.Config with very permissive policy.
func DefaultCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
}
