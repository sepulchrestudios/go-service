package config

import (
	"errors"
	"fmt"
	"sync"

	"github.com/joho/godotenv"
)

// ErrFailedLoadingConfigurationFile is a sentinel error representing a situation where the configuration file could
// not be loaded for some reason.
var ErrFailedLoadingConfigurationFile = errors.New("failed loading configuration file")

// PropertyName represents the name of a configuration key.
type PropertyName string

const (
	// PropertyNameDebugMode represents whether debugging mode is turned on.
	PropertyNameDebugMode PropertyName = "DEBUG"

	// PropertyNameEnvironment represents the environment on which the service is running.
	PropertyNameEnvironment PropertyName = "ENVIRONMENT"

	// PropertyNameGRPCPort represents the port on which the gRPC server will be listening.
	PropertyNameGRPCPort PropertyName = "GRPC_PORT"

	// PropertyNameHttpPort represents the port on which the HTTP service will be listening.
	PropertyNameHTTPPort PropertyName = "PORT"

	// PropertyNameServiceName represents the human-readable name of the service that is running.
	PropertyNameServiceName PropertyName = "NAME"
)

// LoadConfiguration loads the environment configuration from the default path. Returns the config struct instance
// plus any error that may have occurred.
func LoadConfiguration() (*Config, error) {
	return LoadConfigurationFromFile(nil)
}

// LoadConfigurationFromFile loads the environment configuration from the specified file or from the default path if
// the path is nil. Returns the config struct instance plus any error that may have occurred.
func LoadConfigurationFromFile(path *string) (*Config, error) {
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

	config := &Config{
		configuration: configMap,
	}
	return config, nil
}

// NewConfiguration returns a new empty configuration struct instance.
func NewConfiguration() *Config {
	return &Config{
		configuration: map[string]string{},
	}
}

// Config represents a set of mapped configuration values. It also contains a mutex so it should ONLY be passed
// around by-reference and never by-value.
type Config struct {
	configuration map[string]string
	mu            sync.Mutex
}

// GetAllProperties returns a copy of the map that represents all configuration values.
func (c *Config) GetAllProperties() map[string]string {
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
func (c *Config) GetProperty(property PropertyName) (string, bool) {
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
func (c *Config) HasProperty(property PropertyName) bool {
	_, exists := c.GetProperty(property)
	return exists
}

// SetProperty sets the property with the given name to the provided value. Returns a boolean describing whether the
// property was already present and was therefore overwritten.
func (c *Config) SetProperty(property PropertyName, value string) bool {
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
