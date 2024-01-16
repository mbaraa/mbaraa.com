package config

import (
	"os"

	"mbaraacom/log"
)

var (
	_config = config{}
)

func init() {
	_config = config{
		Port:  getEnv("WEBSITE_PORT", "8080"),
		DbUri: getEnv("WEBSITE_DB_URI", ""),
	}
}

type config struct {
	Port  string
	DbUri string
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
