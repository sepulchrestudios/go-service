package feature

import (
	"context"
	"fmt"

	"github.com/open-feature/go-sdk/openfeature"
)

// RegisterOpenFeatureProvider registers the OpenFeature provider with a domain, returning an OpenFeature client, a
// channel that will be closed when the provider is ready, and any error encountered during registration.
func RegisterOpenFeatureProvider(
	ctx context.Context, domain DomainName, provider openfeature.FeatureProvider,
) (*openfeature.Client, chan struct{}, error) {
	if provider == nil {
		return nil, nil, ErrNilFeatureFlagProvider
	}
	if domain == "" {
		return nil, nil, ErrDomainCannotBeEmpty
	}
	domainAsString := string(domain)

	// Set up a channel to signal when the provider is ready since we will be doing asynchronous registration.
	readyChan := make(chan struct{})
	readyCallbackFunc := func(details openfeature.EventDetails) {
		close(readyChan)
	}

	// Register the provider via OpenFeature and add the ready callback handler to it.
	err := openfeature.SetNamedProviderWithContext(ctx, domainAsString, provider)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrFailedToRegisterFeatureFlagProvider, err)
	}
	client := openfeature.NewClient(domainAsString)
	client.AddHandler(openfeature.ProviderReady, &readyCallbackFunc)
	return client, readyChan, nil
}
