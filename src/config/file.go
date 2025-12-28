package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// FileBasedConfig represents a set of mapped configuration values. It also contains a mutex so it should ONLY be
// passed around by-reference and never by-value.
//
// This configuration implementation is intended to load values from a file (e.g., .env file).
type FileBasedConfig struct {
	configuration *ConfigurationMap
}

// GetAllProperties returns a copy of the map that represents all configuration values.
func (c *FileBasedConfig) GetAllProperties() map[string]string {
	if c == nil || c.configuration == nil {
		return map[string]string{}
	}
	return c.configuration.GetAllProperties()
}

// GetProperty takes a property name and returns the matching string value from the configuration plus a boolean
// describing whether the property name actually exists and was therefore valid.
func (c *FileBasedConfig) GetProperty(property PropertyName) (string, bool) {
	if c == nil || c.configuration == nil {
		return "", false
	}
	return c.configuration.GetProperty(property)
}

// HasProperty returns whether the property name exists within the configuration.
func (c *FileBasedConfig) HasProperty(property PropertyName) bool {
	if c == nil || c.configuration == nil {
		return false
	}
	return c.configuration.HasProperty(property)
}

// SetProperty sets the property with the given name to the provided value. Returns a boolean describing whether the
// property was already present and was therefore overwritten.
func (c *FileBasedConfig) SetProperty(property PropertyName, value string) bool {
	if c == nil || c.configuration == nil {
		return false
	}
	return c.configuration.SetProperty(property, value)
}

// NewFileConfiguration returns a new empty file configuration struct instance.
func NewFileConfiguration() *FileBasedConfig {
	return &FileBasedConfig{
		configuration: NewConfigurationMap(),
	}
}

// LoadFileConfiguration loads the environment configuration from the default path. Returns the config struct instance
// plus any error that may have occurred.
func LoadFileConfiguration() (*FileBasedConfig, error) {
	return LoadFileConfigurationFromFile(nil)
}

// LoadFileConfigurationFromFile loads the environment configuration from the specified file or from the default path
// if the path is nil. Returns the config struct instance plus any error that may have occurred.
func LoadFileConfigurationFromFile(path *string) (*FileBasedConfig, error) {
	var err error
	configMap := map[string]string{}

	configFilenames := []string{}
	if path != nil {
		configFilenames = append(configFilenames, *path)
	}

	configMap, err = godotenv.Read(configFilenames...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedLoadingConfigurationFile, err)
	}

	config := &FileBasedConfig{
		configuration: NewConfigurationMapFromMap(configMap),
	}
	return config, nil
}
