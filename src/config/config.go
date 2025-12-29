package config

// Contract is an interface that represents a configuration source for the application.
type Contract interface {
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
