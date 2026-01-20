package main

import (
	"context"
	"fmt"
	"log"
	"net"
	gohttp "net/http"
	"os"
	"strconv"
	"time"

	devcycle "github.com/devcyclehq/go-server-sdk/v2"
	devcycleapi "github.com/devcyclehq/go-server-sdk/v2/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/open-feature/go-sdk/openfeature"
	"github.com/sepulchrestudios/go-service/src/cache"
	"github.com/sepulchrestudios/go-service/src/config"
	"github.com/sepulchrestudios/go-service/src/database"
	"github.com/sepulchrestudios/go-service/src/event"
	"github.com/sepulchrestudios/go-service/src/feature"
	servicelogger "github.com/sepulchrestudios/go-service/src/log"
	"github.com/sepulchrestudios/go-service/src/mail"
	"github.com/sepulchrestudios/go-service/src/server"
	"github.com/sepulchrestudios/go-service/src/service"
	"github.com/sepulchrestudios/go-service/src/work"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Connect to the intended cache using the provided environment configuration. Returns the cache implementation plus
// any error that may have occurred.
func connectToCacheFromConfig(
	ctx context.Context, envConfig config.Contract, isDebugModeActive bool, debugLogger servicelogger.DebugContract,
) (cache.Contract, error) {
	// Resolve the cache password from either a file path or the direct property
	//
	// Cache password is optional so being completely blank regardless of existence is a valid scenario
	cachePassword, _, err := readSecretFromFileOrEnvFallback(
		config.PropertyNameCachePasswordFile, config.PropertyNameCachePassword, envConfig,
	)
	if err != nil {
		return nil, err
	}

	// Create the cache connection from the environment configuration
	var cacheAddr string
	cacheHost, exists := envConfig.GetProperty(config.PropertyNameCacheHost)
	if !exists {
		return nil, fmt.Errorf("Cannot read property from configuration: %s", config.PropertyNameCacheHost)
	}
	cachePort, _ := envConfig.GetProperty(config.PropertyNameCachePort)
	if cachePort != "" {
		cacheAddr = fmt.Sprintf("%s:%s", cacheHost, cachePort)
	} else {
		cacheAddr = cacheHost
	}
	cacheIdentifier, exists := envConfig.GetProperty(config.PropertyNameCacheIdentifier)
	if !exists {
		return nil, fmt.Errorf("Cannot read property from configuration: %s", config.PropertyNameCacheIdentifier)
	}
	cacheUsername, _ := envConfig.GetProperty(config.PropertyNameCacheUsername)
	connectionArguments := &cache.RedisConnectionArguments{
		CacheConnectionArguments: cache.CacheConnectionArguments{
			CacheIdentifier: cacheIdentifier,
		},
		Addr:     cacheAddr,
		Password: cachePassword,
		Username: cacheUsername,
	}
	cacheImplementation, err := cache.NewRedis(ctx, connectionArguments)
	if err != nil {
		return nil, err
	}
	if isDebugModeActive {
		return cache.NewDebug(cacheImplementation, debugLogger), nil
	}
	return cacheImplementation, nil
}

// Connect to the intended database using the provided environment configuration. Returns the database connection plus
// any error that may have occurred.
func connectToDatabaseFromConfig(
	envConfig config.Contract, isDebugModeActive bool,
) (database.Contract, error) {
	// Resolve the DB password from either a file path or the direct property
	dbPassword, exists, err := readSecretFromFileOrEnvFallback(
		config.PropertyNameDatabasePasswordFile, config.PropertyNameDatabasePassword, envConfig,
	)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("Cannot read property from configuration (make sure it exists): %s",
			config.PropertyNameDatabasePassword)
	}

	// Create the Postgres connection from the environment configuration
	dbHost, exists := envConfig.GetProperty(config.PropertyNameDatabaseHost)
	if !exists {
		return nil, fmt.Errorf("Cannot read property from configuration: %s", config.PropertyNameDatabaseHost)
	}
	dbUsername, exists := envConfig.GetProperty(config.PropertyNameDatabaseUsername)
	if !exists {
		return nil, fmt.Errorf("Cannot read property from configuration: %s", config.PropertyNameDatabaseUsername)
	}
	dbName, exists := envConfig.GetProperty(config.PropertyNameDatabaseName)
	if !exists {
		return nil, fmt.Errorf("Cannot read property from configuration: %s", config.PropertyNameDatabaseName)
	}
	dbPort, _ := envConfig.GetProperty(config.PropertyNameDatabasePort)
	dbSSLMode, _ := envConfig.GetProperty(config.PropertyNameDatabaseSSLMode)
	dbTimezone, _ := envConfig.GetProperty(config.PropertyNameDatabaseTimezone)
	connectionArguments := &database.PostgresDatabaseConnectionArguments{
		DatabaseConnectionArguments: database.DatabaseConnectionArguments{
			DatabaseName: dbName,
			Host:         dbHost,
			Password:     dbPassword,
			Port:         dbPort,
			Username:     dbUsername,
		},
		SSLMode:  dbSSLMode,
		Timezone: dbTimezone,
	}
	return database.NewPostgresDatabaseConnection(connectionArguments, isDebugModeActive)
}

// Connect to the intended feature flag service using the provided environment configuration. Returns the OpenFeature
// provider, a channel that will be unblocked upon readiness, plus any error that may have occurred.
func connectToFeatureFlagServiceFromConfig(
	envConfig config.Contract,
) (openfeature.FeatureProvider, chan devcycleapi.ClientEvent, error) {
	// Resolve the SDK key from either a file path or the direct property
	sdkKey, exists, err := readSecretFromFileOrEnvFallback(
		config.PropertyNameFeatureFlagSDKKeyFile, config.PropertyNameFeatureFlagSDKKey, envConfig,
	)
	if err != nil {
		return nil, nil, err
	}
	if !exists {
		return nil, nil, fmt.Errorf("Cannot read property from configuration (make sure it exists): %s",
			config.PropertyNameFeatureFlagSDKKey)
	}
	// Create the DevCycle client and OpenFeature provider with basic options
	onInitializedChannel := make(chan devcycleapi.ClientEvent)
	options := devcycle.Options{
		ClientEventHandler:           onInitializedChannel,
		EnableEdgeDB:                 false,
		EnableCloudBucketing:         false,
		EventFlushIntervalMS:         30 * time.Second,
		ConfigPollingIntervalMS:      1 * time.Minute,
		RequestTimeout:               30 * time.Second,
		DisableAutomaticEventLogging: false,
		DisableCustomEventLogging:    false,
	}
	_, provider, err := feature.NewDevCycleClient(sdkKey, &options)
	if err != nil {
		return nil, nil, err
	}
	return provider, onInitializedChannel, nil
}

// pumpEventBus pumps events from the provided event bus in its own goroutine.
func pumpEventBus(ctx context.Context, eventBus work.BusPumperContract, debugLogger servicelogger.DebugContract) {
	go func(ctx context.Context, bus work.BusPumperContract, logger servicelogger.DebugContract) {
		logger.Debug("Starting event bus pump...")
		if bus == nil {
			logger.Debug("No event bus instance provided; skipping event bus pump.")
			return
		}
		err := bus.Pump(ctx)
		logger.Debug("Finished pumping events.", zap.Error(err))
	}(ctx, eventBus, debugLogger)
}

// pumpMailBus pumps events from the provided mail bus in its own goroutine.
func pumpMailBus(ctx context.Context, mailBus work.BusPumperContract, debugLogger servicelogger.DebugContract) {
	go func(ctx context.Context, bus work.BusPumperContract, logger servicelogger.DebugContract) {
		logger.Debug("Starting mail bus pump...")
		if bus == nil {
			logger.Debug("No mail bus instance provided; skipping mail bus pump.")
			return
		}
		err := bus.Pump(ctx)
		logger.Debug("Finished pumping mail messages.", zap.Error(err))
	}(ctx, mailBus, debugLogger)
}

// readSecretFromFileOrEnvFallback attempts to read a secret from a file path specified in the configuration. If the
// file path property does not exist or is otherwise empty, it falls back to reading the secret directly from the
// environment configuration. It returns the resolved secret, a boolean indicating whether the secret was found, and
// any error encountered during the process.
func readSecretFromFileOrEnvFallback(
	filePathPropertyName config.PropertyName, envPropertyName config.PropertyName, envConfig config.Contract,
) (string, bool, error) {
	// Resolve the secret from either a file path or the direct property
	var secret string
	secretFilePath, exists := envConfig.GetProperty(filePathPropertyName)
	if exists && secretFilePath != "" {
		secretBytes, err := os.ReadFile(secretFilePath)
		if err != nil {
			return "", exists, fmt.Errorf("Cannot read secret from file path '%s': %v", secretFilePath, err)
		}
		secret = string(secretBytes)
	} else {
		// Secret being completely blank regardless of existence is also a valid scenario
		secret, exists = envConfig.GetProperty(envPropertyName)
	}
	return secret, exists, nil
}

func main() {
	var envConfig config.Contract
	var err error

	// Determine whether to load environment configuration from a file or directly from environment variables
	shouldLoadFromFileStr, exists := os.LookupEnv(string(config.PropertyNameLoadEnvFromFile))
	shouldLoadFromFile, _ := strconv.ParseBool(shouldLoadFromFileStr)
	if exists && !shouldLoadFromFile {
		envConfig, err = config.LoadEnvironmentConfiguration()
		if err != nil {
			err = fmt.Errorf("%w: [environment-based]", err)
		}
	} else {
		envConfig, err = config.LoadFileConfiguration()
		if err != nil {
			err = fmt.Errorf("%w: [file-based]", err)
		}
	}
	if err != nil {
		log.Fatalln("Cannot load environment configuration: ", err)
	}

	// Create the standard logger
	isDebugModeActiveStr, _ := envConfig.GetProperty(config.PropertyNameDebugMode)
	isDebugModeActive, _ := strconv.ParseBool(isDebugModeActiveStr)
	logger, err := servicelogger.NewStandardLogger(isDebugModeActive)
	if err != nil {
		log.Fatalln("Cannot create standard logger: ", err)
	}

	// Create the database connection here
	logger.Info("Connecting to database...")
	_, err = connectToDatabaseFromConfig(envConfig, isDebugModeActive)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cannot connect to database: %v", err))
	}
	logger.Info("Connected to database successfully")

	// Create the cache connection here
	logger.Info("Connecting to cache...")
	ctx := context.Background()
	_, err = connectToCacheFromConfig(ctx, envConfig, isDebugModeActive, logger)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cannot connect to cache: %v", err))
	}
	logger.Info("Connected to cache successfully")

	// Create a cancellable context for the event and mail bus processors
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start the event bus processor with a single registered default handler
	eventBus := event.NewBus(work.NewConcurrentBus())
	err = eventBus.RegisterDefaultHandler()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cannot register default event handler: %v", err))
	}
	pumpEventBus(cancelCtx, eventBus, logger)

	// Start the mail bus processor with a single registered default handler
	mailBus := mail.NewBus(work.NewConcurrentBus())
	err = mailBus.RegisterDefaultHandler()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cannot register default message handler: %v", err))
	}
	pumpMailBus(cancelCtx, mailBus, logger)

	// Set up the feature flag provider to work with OpenFeature; in our case, we're using DevCycle
	logger.Info("Setting up feature flag provider...")
	featureFlagProvider, devCycleReadyChan, err := connectToFeatureFlagServiceFromConfig(envConfig)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cannot set up feature flag provider from DevCycle client: %v", err))
	}
	_, openFeatureReadyChan, err := feature.RegisterOpenFeatureProvider(
		ctx, feature.DomainNameFeatureFlags, featureFlagProvider,
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cannot register feature flag provider with OpenFeature: %v", err))
	}
	logger.Info("Feature flag provider set up successfully")

	// Resolve the necessary ports
	grpcPort, exists := envConfig.GetProperty(config.PropertyNameGRPCPort)
	if !exists {
		logger.Fatal(fmt.Sprintf("Cannot read property from configuration: %s", config.PropertyNameGRPCPort))
	}
	httpPort, exists := envConfig.GetProperty(config.PropertyNameHTTPPort)
	if !exists {
		logger.Fatal(fmt.Sprintf("Cannot read property from configuration: %s", config.PropertyNameHTTPPort))
	}

	// Much of this HTTP/gRPC stuff comes from the barebones gRPC Gateway functionality:
	// https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen: %v", err))
	}

	// Create a gRPC server object and liveness server object, then attach them
	grpcServer := grpc.NewServer()
	livenessServer := server.NewLivenessServer(service.NewLivenessService())
	server.RegisterLivenessServer(grpcServer, livenessServer)

	// Serve gRPC server
	logger.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%s", grpcPort))
	go func() {
		logger.Fatal(grpcServer.Serve(lis).Error())
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%s", grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to dial server: %v", err))
	}

	// Register liveness server with the mux and client connection
	gwmux := runtime.NewServeMux()
	err = server.RegisterLivenessServerHandlers(context.Background(), gwmux, conn)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to register gateway: %v", err))
	}

	gwServer := &gohttp.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: gwmux,
	}

	logger.Info(fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%s", httpPort))
	logger.Fatal(func() string {
		// Wait for the feature flag provider to be initialized before serving traffic
		<-devCycleReadyChan
		logger.Info("DevCycle client initialized")
		<-openFeatureReadyChan
		logger.Info("OpenFeature provider initialized")
		// Signal that we are ready to receive traffic and then serve the endpoints
		err := livenessServer.MarkReady()
		if err != nil {
			return err.Error()
		}
		logger.Info("Service is now marked as ready to receive traffic")
		err = gwServer.ListenAndServe()
		if err != nil {
			return err.Error()
		}
		return "Finished"
	}())
}
