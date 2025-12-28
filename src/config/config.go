package config

import "sync"

// PropertyName represents the name of a configuration key.
type PropertyName string

const (
	// PropertyNameCacheHost represents the cache host address.
	PropertyNameCacheHost PropertyName = "CACHE_HOST"

	// PropertyNameCachePort represents the cache port.
	PropertyNameCachePort PropertyName = "CACHE_PORT"

	// PropertyNameCacheName represents the cache name.
	PropertyNameCacheName PropertyName = "CACHE_NAME"

	// PropertyNameCacheUsername represents the cache username.
	PropertyNameCacheUsername PropertyName = "CACHE_USERNAME"

	// PropertyNameCachePassword represents the cache password.
	PropertyNameCachePassword PropertyName = "CACHE_PASSWORD"

	// PropertyNameCachePasswordFile represents the file path from which to read the cache password.
	PropertyNameCachePasswordFile PropertyName = "CACHE_PASSWORD_FILE"

	// PropertyNameDatabaseHost represents the database host address.
	PropertyNameDatabaseHost PropertyName = "DATABASE_HOST"

	// PropertyNameDatabaseName represents the database name.
	PropertyNameDatabaseName PropertyName = "DATABASE_NAME"

	// PropertyNameDatabasePassword represents the database password.
	PropertyNameDatabasePassword PropertyName = "DATABASE_PASSWORD"

	// PropertyNameDatabasePasswordFile represents the file path from which to read the database password.
	PropertyNameDatabasePasswordFile PropertyName = "DATABASE_PASSWORD_FILE"

	// PropertyNameDatabasePort represents the database port.
	PropertyNameDatabasePort PropertyName = "DATABASE_PORT"

	// PropertyNameDatabaseUsername represents the database username.
	PropertyNameDatabaseUsername PropertyName = "DATABASE_USERNAME"

	// PropertyNameDatabaseSSLMode represents the database SSL mode configuration.
	PropertyNameDatabaseSSLMode PropertyName = "DATABASE_SSL_MODE"

	// PropertyNameDatabaseTimezone represents the database default timezone.
	PropertyNameDatabaseTimezone PropertyName = "DATABASE_TIMEZONE"

	// PropertyNameDebugMode represents whether debugging mode is turned on.
	PropertyNameDebugMode PropertyName = "DEBUG"

	// PropertyNameEnvironment represents the environment on which the service is running.
	PropertyNameEnvironment PropertyName = "ENVIRONMENT"

	// PropertyNameGRPCPort represents the port on which the gRPC server will be listening.
	PropertyNameGRPCPort PropertyName = "GRPC_PORT"

	// PropertyNameHttpPort represents the port on which the HTTP service will be listening.
	PropertyNameHTTPPort PropertyName = "PORT"

	// PropertyNameLoadEnvFromFile represents whether to load environment variables from a .env file.
	PropertyNameLoadEnvFromFile PropertyName = "LOAD_ENV_FROM_FILE"

	// PropertyNameServiceName represents the human-readable name of the service that is running.
	PropertyNameServiceName PropertyName = "NAME"
)

// GetAvailableConfigurationKeys returns a slice of all available configuration property names.
func GetAvailableConfigurationKeys() []PropertyName {
	return []PropertyName{
		PropertyNameCacheHost,
		PropertyNameCachePort,
		PropertyNameCacheName,
		PropertyNameCacheUsername,
		PropertyNameCachePassword,
		PropertyNameCachePasswordFile,
		PropertyNameDatabaseHost,
		PropertyNameDatabaseName,
		PropertyNameDatabasePassword,
		PropertyNameDatabasePasswordFile,
		PropertyNameDatabasePort,
		PropertyNameDatabaseUsername,
		PropertyNameDatabaseSSLMode,
		PropertyNameDatabaseTimezone,
		PropertyNameDebugMode,
		PropertyNameEnvironment,
		PropertyNameGRPCPort,
		PropertyNameHTTPPort,
		PropertyNameLoadEnvFromFile,
		PropertyNameServiceName,
	}
}

// Config is an interface that represents a configuration source for the application.
type Config interface {
	// GetAllProperties returns a copy of the map that represents all configuration values.
	GetAllProperties() map[string]string

	// GetProperty takes a property name and returns the matching string value from the configuration plus a boolean
	// describing whether the property name actually exists and was therefore valid.
	GetProperty(property PropertyName) (string, bool)

	// HasProperty returns whether the property name exists within the configuration.
	HasProperty(property PropertyName) bool

	// SetProperty sets the property with the given name to the provided value. Returns a boolean describing whether the
	// property was already present and was therefore overwritten.
	SetProperty(property PropertyName, value string) bool
}

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
