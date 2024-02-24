package config

import (
	"os"

	"github.com/mbaraa/mbaraa.com/log"
)

var (
	_config = config{}
)

func init() {
	_config = config{
		WebsitePort: getEnv("WEBSITE_PORT", "8080"),
		FilesDir:    getEnv("FILES_DIR", "./_files/"),
	}
}

type config struct {
	WebsitePort string
	FilesDir    string
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
