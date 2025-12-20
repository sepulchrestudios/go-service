package config

import "os"

// EnvironmentBasedConfig represents a set of mapped configuration values. It also contains a mutex so it should ONLY
// be passed around by-reference and never by-value.
//
// This configuration implementation is intended to load values from environment variables.
type EnvironmentBasedConfig struct {
	configuration *ConfigurationMap
}

// GetAllProperties returns a copy of the map that represents all configuration values.
func (c *EnvironmentBasedConfig) GetAllProperties() map[string]string {
	if c == nil || c.configuration == nil {
		return map[string]string{}
	}
	return c.configuration.GetAllProperties()
}

// GetProperty takes a property name and returns the matching string value from the configuration plus a boolean
// describing whether the property name actually exists and was therefore valid.
func (c *EnvironmentBasedConfig) GetProperty(property PropertyName) (string, bool) {
	if c == nil || c.configuration == nil {
		return "", false
	}
	return c.configuration.GetProperty(property)
}

// HasProperty returns whether the property name exists within the configuration.
func (c *EnvironmentBasedConfig) HasProperty(property PropertyName) bool {
	if c == nil || c.configuration == nil {
		return false
	}
	return c.configuration.HasProperty(property)
}

// SetProperty sets the property with the given name to the provided value. Returns a boolean describing whether the
// property was already present and was therefore overwritten.
func (c *EnvironmentBasedConfig) SetProperty(property PropertyName, value string) bool {
	if c == nil || c.configuration == nil {
		return false
	}
	return c.configuration.SetProperty(property, value)
}

// NewEnvironmentConfiguration returns a new empty environment configuration struct instance.
func NewEnvironmentConfiguration() *EnvironmentBasedConfig {
	return &EnvironmentBasedConfig{
		configuration: NewConfigurationMap(),
	}
}

// LoadEnvironmentConfiguration loads the environment configuration from the current environment variables. Returns the
// config struct instance plus any error that may have occurred.
func LoadEnvironmentConfiguration() (*EnvironmentBasedConfig, error) {
	configMap := map[string]string{}

	availableKeys := GetAvailableConfigurationKeys()
	for _, key := range availableKeys {
		value, exists := os.LookupEnv(string(key))
		if exists {
			configMap[string(key)] = value
		}
	}

	config := &EnvironmentBasedConfig{
		configuration: NewConfigurationMapFromMap(configMap),
	}
	return config, nil
}
