package config

import "errors"

// ErrFailedLoadingConfigurationFile is a sentinel error representing a situation where the configuration file could
// not be loaded for some reason.
var ErrFailedLoadingConfigurationFile = errors.New("failed loading configuration file")
