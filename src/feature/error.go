package feature

import "errors"

// ErrDomainCannotBeEmpty is a sentinel error that represents an empty domain name being supplied.
var ErrDomainCannotBeEmpty = errors.New("domain cannot be empty")

// ErrFailedToInitializeFeatureFlagClient is a sentinel error that represents a failure when initializing a feature
// flag client.
var ErrFailedToInitializeFeatureFlagClient = errors.New("failed to initialize feature flag client")

// ErrFailedToRegisterFeatureFlagProvider is a sentinel error that represents a failure when registering a feature
// flag provider.
var ErrFailedToRegisterFeatureFlagProvider = errors.New("failed to register feature flag provider")

// ErrNilDevCycleClient is a sentinel error that represents a nil DevCycle client being supplied.
var ErrNilDevCycleClient = errors.New("DevCycle client is nil")

// ErrNilFeatureFlagClient is a sentinel error that represents a nil feature flag client being supplied.
var ErrNilFeatureFlagClient = errors.New("feature flag client cannot be nil")

// ErrNilFeatureFlagProvider is a sentinel error that represents a nil feature flag provider being supplied.
var ErrNilFeatureFlagProvider = errors.New("feature flag provider cannot be nil")
