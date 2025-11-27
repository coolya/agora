package config

import "github.com/spf13/viper"

// Config holds the entire application configuration
type Config struct {
	Sources []Source `mapstructure:"sources"`
}

// Source represents a single source of ADRs, like a GitHub repository
type Source struct {
	Type   string `mapstructure:"type"`
	URL    string `mapstructure:"url"`
	Path   string `mapstructure:"path"`
	Auth   Auth   `mapstructure:"auth"`
	GitLab GitLab `mapstructure:"gitlab"`
}

// Auth holds the authentication details for a source
type Auth struct {
	Token string `mapstructure:"token"`
}

// GitLab holds GitLab-specific configuration.
type GitLab struct {
	BaseURL string `mapstructure:"baseURL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
