package config

import "sync"

// ConfigurationMap represents a set of mapped configuration values. It also contains a mutex so it should ONLY be
// passed around by-reference and never by-value.
type ConfigurationMap struct {
	configuration map[string]string
	mu            sync.Mutex
}

// GetAllProperties returns a copy of the map that represents all configuration values.
func (c *ConfigurationMap) GetAllProperties() map[string]string {
	allConfigProperties := map[string]string{}
	if c == nil || c.configuration == nil {
		return allConfigProperties
	}

	// ensure we don't get a collision if two or more goroutines try to read concurrently
	c.mu.Lock()
	defer c.mu.Unlock()
	for configKey, configValue := range c.configuration {
		allConfigProperties[configKey] = configValue
	}
	return allConfigProperties
}

// GetProperty takes a property name and returns the matching string value from the configuration plus a boolean
// describing whether the property name actually exists and was therefore valid.
func (c *ConfigurationMap) GetProperty(property PropertyName) (string, bool) {
	if c == nil || c.configuration == nil {
		return "", false
	}

	// ensure we don't get a collision if two or more goroutines try to read concurrently
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.configuration[string(property)]
	return value, exists
}

// HasProperty returns whether the property name exists within the configuration.
func (c *ConfigurationMap) HasProperty(property PropertyName) bool {
	if c == nil || c.configuration == nil {
		return false
	}
	_, exists := c.GetProperty(property)
	return exists
}

// SetProperty sets the property with the given name to the provided value. Returns a boolean describing whether the
// property was already present and was therefore overwritten.
func (c *ConfigurationMap) SetProperty(property PropertyName, value string) bool {
	if c == nil || c.configuration == nil {
		return false
	}
	propertyKey := string(property)

	// ensure we don't get a collision if two or more goroutines try to write concurrently
	c.mu.Lock()
	defer c.mu.Unlock()
	_, willBeOverwritten := c.configuration[propertyKey]
	c.configuration[propertyKey] = value
	return willBeOverwritten
}

// NewConfigurationMap returns a new empty configuration map struct instance.
func NewConfigurationMap() *ConfigurationMap {
	return &ConfigurationMap{
		configuration: map[string]string{},
	}
}

// NewConfigurationMapFromMap returns a new configuration map struct instance initialized with the provided map.
func NewConfigurationMapFromMap(input map[string]string) *ConfigurationMap {
	configMap := NewConfigurationMap()
	for key, val := range input {
		configMap.configuration[key] = val
	}
	return configMap
}
