package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config holds the entire application configuration
type Config struct {
	Sources []Source `mapstructure:"sources"`
}

// Source represents a single source of ADRs, like a GitHub repository
type Source struct {
	Name   string `mapstructure:"name"`
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
	if err != nil {
		return
	}

	// Auto-generate names for sources without names and validate uniqueness
	err = config.ensureUniqueSourceNames()
	return
}

// ensureUniqueSourceNames auto-generates names for sources without names and validates uniqueness
func (c *Config) ensureUniqueSourceNames() error {
	seenNames := make(map[string]bool)
	typeCounters := make(map[string]int)

	for i := range c.Sources {
		// If name is already set, just validate and record it
		if c.Sources[i].Name != "" {
			if seenNames[c.Sources[i].Name] {
				return fmt.Errorf("duplicate source name: %s", c.Sources[i].Name)
			}
			seenNames[c.Sources[i].Name] = true
			continue
		}

		// Name is empty: generate one based on type and counter, avoiding collisions
		sourceType := c.Sources[i].Type
		for {
			counter := typeCounters[sourceType]
			candidate := fmt.Sprintf("%s-%d", sourceType, counter)
			typeCounters[sourceType] = counter + 1

			if !seenNames[candidate] {
				c.Sources[i].Name = candidate
				seenNames[candidate] = true
				break
			}
		}
	}

	return nil
}

