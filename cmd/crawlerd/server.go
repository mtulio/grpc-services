package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/mtulio/grpc-services/src/crawler"
)

func main() {
	serverPort, err := strconv.Atoi(os.Getenv("FIM_CRAWLER_GRPC_SERVER_PORT"))
	if err != nil {
		serverPort = 50051
	}
	if serverPort == 0 {
		serverPort = 50051
	}
	serverAddr := fmt.Sprintf("0.0.0.0:%d", serverPort)
	fmt.Println("Starting gRPC server on %s", serverAddr)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	tls := false
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	crawler.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
