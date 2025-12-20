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
	"github.com/sepulchrestudios/go-service/src/config"
	servicelogger "github.com/sepulchrestudios/go-service/src/log"
	"github.com/sepulchrestudios/go-service/src/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Much of this comes from the barebones gRPC Gateway functionality:
	// https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc

	var envConfig config.Config
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
	livenessServer := server.NewLivenessServer()
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
	logger.Fatal(gwServer.ListenAndServe().Error())
}
