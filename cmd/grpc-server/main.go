package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/mtulio/grpc-services/src/crawler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	serverPort, err := strconv.Atoi(os.Getenv("GRPC_SERVER_PORT"))
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
	// pb.RegisterWebCrawlerServiceServer(s, &server{})
	crawler.Register(s)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer s.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	log.Println("signal received")
}
