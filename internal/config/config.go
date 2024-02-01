package config

import (
	"os"
	"time"

	"internal/log"
)

var (
	_config = config{}
)

func init() {
	_config = config{
		WebsitePort:       getEnv("WEBSITE_PORT", "8080"),
		DashboardPort:     getEnv("DASHBOARD_PORT", "8081"),
		DashboardPassword: getEnv("DASHBOARD_PASSWORD", time.Now().String()),
		DbUri:             getEnv("DB_URI", ""),
	}
}

type config struct {
	WebsitePort       string
	DashboardPort     string
	DashboardPassword string
	DbUri             string
}

// Config returns the API's config :)
func Config() config {
	return _config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Warningf("The \"%s\" variable is not set. Defaulting to \"%s\".\n", key, defaultValue)
		value = defaultValue
	}
	return value
}
