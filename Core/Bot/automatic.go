package Bot

import (
	"BoostTool/Core/Utils"
	"log"

	"github.com/gin-gonic/gin"
)

// ServerConfig holds the configuration for the server
type ServerConfig struct {
	Port string
}

// Automation sets up and starts the Gin server
func Automation() {
	config := loadServerConfig() // Assume this function loads the server configuration

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Start the server
	Utils.LogInfo("Running on port", "Port", config.Port)
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// loadServerConfig loads the server configuration
func loadServerConfig() *ServerConfig {
	// This should ideally load the configuration from a file or environment variables
	return &ServerConfig{
		Port: "8080", // Default port
	}
}
