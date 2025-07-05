package main

import (
	"context"
	"fmt"
	"log"
	"net"
	gohttp "net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sepulchrestudios/go-service/src/config"
	"github.com/sepulchrestudios/go-service/src/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Much of this comes from the barebones gRPC Gateway functionality:
	// https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#using-protoc

	// Load the environment configuration
	envConfig, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalln("Cannot load configuration: ", err)
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
	log.Println(fmt.Sprintf("Serving gRPC on 0.0.0.0:%s", grpcPort))
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

	log.Println(fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%s", httpPort))
	log.Fatalln(gwServer.ListenAndServe())
}
