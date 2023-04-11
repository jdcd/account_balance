package config

import (
	"os"
)

// AppConfiguration contains data from env variables, to works inside the app
type AppConfiguration struct {
	RepositoryURL    string
	EmailDestination string
	EmailApiKey      string
}

// GetConfigurations return *AppConfiguration instance, with configured data
func GetConfigurations() *AppConfiguration {
	return &AppConfiguration{
		RepositoryURL:    os.Getenv("CONNECTION_URL"),
		EmailDestination: os.Getenv("EMAIL_DESTINATION"),
		EmailApiKey:      os.Getenv("EMAIL_API_KEY"),
	}
}
