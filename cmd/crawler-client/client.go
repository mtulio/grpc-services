package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc/credentials"

	pb "github.com/mtulio/grpc-services/src/protobuf/crawler"

	"google.golang.org/grpc"
)

func getServerAddress() string {
	serverPort, err := strconv.Atoi(os.Getenv("FIM_CRAWLER_GRPC_SERVER_PORT"))
	if err != nil {
		serverPort = 50051
	}
	if serverPort == 0 {
		serverPort = 50051
	}

	serverHost := os.Getenv("FIM_CRAWLER_GRPC_SERVER_HOST")
	if serverHost == "" {
		serverHost = "localhost"
	}
	return fmt.Sprintf("%s:%d", serverHost, serverPort)

}

func main() {
	serverAddr := getServerAddress()
	fmt.Println("Hello I'm a client")
	fmt.Println(" I will connect to the server %s", serverAddr)

	option := flag.String("function", "unary", "RPC functions: unary, bidi")
	flag.Parse()

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial(serverAddr, opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewWebCrawlerServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	switch *option {
	case "bidi":
		doBiDiStreaming(c)
	case "unary":
		doUnary(c)
	default:
		doUnary(c)
	}

}

func doUnary(c pb.WebCrawlerServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &pb.WebCrawlerRequest{
		Name: "Google",
		Url:  "https://google.com/",
	}
	res, err := c.WebCrawler(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v %v", res.GetStatus(), len(res.GetPayload()))
}

func doBiDiStreaming(c pb.WebCrawlerServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.WebCrawlerBatch(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := &pb.WebCrawlerBatchRequest{
		Requests: []*pb.WebCrawlerBatchRequest_Request{
			{
				Name: "Google",
				Url:  "https://google.com/",
			},
			{
				Name: "Yahoo",
				Url:  "https://Yahoo.com/",
			},
			{
				Name: "a",
				Url:  "https://a.com/",
			},
			{
				Name: "b",
				Url:  "https://b.com/",
			},
			{
				Name: "c",
				Url:  "https://c.com/",
			},
			{
				Name: "d",
				Url:  "https://d.com/",
			},
			{
				Name: "e",
				Url:  "https://e.com/",
			},
			{
				Name: "f",
				Url:  "https://f.com/",
			},
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		stream.Send(requests)
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received ["+res.GetName()+"]: %v\n", res.GetStatus())
		}
		close(waitc)
	}()

	<-waitc
}
