package feature

import (
	"fmt"

	devcycle "github.com/devcyclehq/go-server-sdk/v2"
	"github.com/open-feature/go-sdk/openfeature"
)

// NewDevCycleClient creates a new DevCycle client based upon the passed SDK key and options, returning the DevCycle
// client, the matching OpenFeature feature provider, and any error encountered during creation.
func NewDevCycleClient(
	sdkKey string, options *devcycle.Options,
) (*devcycle.Client, openfeature.FeatureProvider, error) {
	dvcClient, err := devcycle.NewClient(sdkKey, options)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrFailedToInitializeFeatureFlagClient, err)
	}
	if dvcClient == nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrNilFeatureFlagClient, ErrNilDevCycleClient)
	}
	return dvcClient, dvcClient.OpenFeatureProvider(), nil
}
