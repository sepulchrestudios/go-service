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
	if exists && shouldLoadFromFile {
		envConfig, err = config.LoadFileConfiguration()
		if err != nil {
			err = fmt.Errorf("%w: [file-based]", err)
		}
	} else {
		envConfig, err = config.LoadEnvironmentConfiguration()
		if err != nil {
			err = fmt.Errorf("%w: [environment-based]", err)
		}
	}
	if err != nil {
		log.Fatalln("Cannot load environment configuration: ", err)
	}

	// Resolve the necessary ports
	grpcPort, exists := envConfig.GetProperty(config.PropertyNameGRPCPort)
	if !exists {
		log.Fatalln("Cannot read property from configuration: ", config.PropertyNameGRPCPort)
	}
	httpPort, exists := envConfig.GetProperty(config.PropertyNameHTTPPort)
	if !exists {
		log.Fatalln("Cannot read property from configuration: ", config.PropertyNameHTTPPort)
	}

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object and liveness server object, then attach them
	grpcServer := grpc.NewServer()
	livenessServer := server.NewLivenessServer()
	server.RegisterLivenessServer(grpcServer, livenessServer)

	// Serve gRPC server
	log.Printf("Serving gRPC on 0.0.0.0:%s\n", grpcPort)
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%s", grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// Register liveness server with the mux and client connection
	gwmux := runtime.NewServeMux()
	err = server.RegisterLivenessServerHandlers(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &gohttp.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%s\n", httpPort)
	log.Fatalln(gwServer.ListenAndServe())
}
