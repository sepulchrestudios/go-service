package main

import (
	"context"
	"fmt"
	"log"
	"net"
	gohttp "net/http"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sepulchrestudios/go-service/src/cache"
	"github.com/sepulchrestudios/go-service/src/config"
	"github.com/sepulchrestudios/go-service/src/database"
	servicelogger "github.com/sepulchrestudios/go-service/src/log"
	"github.com/sepulchrestudios/go-service/src/server"
	"github.com/sepulchrestudios/go-service/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Connect to the intended cache using the provided environment configuration. Returns the cache implementation plus
// any error that may have occurred.
func connectToCacheFromConfig(
	ctx context.Context, envConfig config.Contract, isDebugModeActive bool, debugLogger servicelogger.DebugContract,
) (cache.Contract, error) {
	// Resolve the cache password from either a file path or the direct property
	var cachePassword string
	cachePasswordFilePath, exists := envConfig.GetProperty(config.PropertyNameCachePasswordFile)
	if exists && cachePasswordFilePath != "" {
		passwordBytes, err := os.ReadFile(cachePasswordFilePath)
		if err != nil {
			return nil, fmt.Errorf("Cannot read cache password from file path '%s': %v", cachePasswordFilePath, err)
		}
		cachePassword = string(passwordBytes)
	} else {
		// Cache password is optional so being completely blank regardless of existence is a valid scenario
		cachePassword, _ = envConfig.GetProperty(config.PropertyNameCachePassword)
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
	var dbPassword string
	dbPasswordFilePath, exists := envConfig.GetProperty(config.PropertyNameDatabasePasswordFile)
	if exists && dbPasswordFilePath != "" {
		passwordBytes, err := os.ReadFile(dbPasswordFilePath)
		if err != nil {
			return nil, fmt.Errorf("Cannot read database password from file path '%s': %v", dbPasswordFilePath, err)
		}
		dbPassword = string(passwordBytes)
	} else {
		dbPassword, exists = envConfig.GetProperty(config.PropertyNameDatabasePassword)
		if !exists {
			return nil, fmt.Errorf("Cannot read property from configuration (make sure it exists): %s",
				config.PropertyNameDatabasePassword)
		}
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

func main() {
	// Much of this comes from the barebones gRPC Gateway functionality:
	// https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc

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

	// Resolve the necessary ports
	grpcPort, exists := envConfig.GetProperty(config.PropertyNameGRPCPort)
	if !exists {
		logger.Fatal(fmt.Sprintf("Cannot read property from configuration: %s", config.PropertyNameGRPCPort))
	}
	httpPort, exists := envConfig.GetProperty(config.PropertyNameHTTPPort)
	if !exists {
		logger.Fatal(fmt.Sprintf("Cannot read property from configuration: %s", config.PropertyNameHTTPPort))
	}

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
		// Signal that we are ready to receive traffic and then serve the endpoints
		err := livenessServer.MarkReady()
		if err != nil {
			return err.Error()
		}
		err = gwServer.ListenAndServe()
		if err != nil {
			return err.Error()
		}
		return "Finished"
	}())
}
