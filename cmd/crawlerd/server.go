package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/mtulio/grpc-services/src/protobuf/pbcrawler"
)

type server struct{}

func (*server) WebCrawler(ctx context.Context, req *pb.WebCrawlerRequest) (*pb.WebCrawlerResponse, error) {
	fmt.Printf("WebCrawler function was invoked with %v\n", req)
	name := req.GetName()
	url := req.GetUrl()
	status := false
	bodyString := ""
	var bodyBytes []byte

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			bodyString = fmt.Sprintf("error, statusCode: %s", response.StatusCode)
		} else {
			bodyString = string(bodyBytes)
			status = true
		}
	} else {
		bodyString = fmt.Sprintf("error, statusCode: %s", response.StatusCode)
	}

	fmt.Println("Body length: ", len(bodyString))
	res := &pb.WebCrawlerResponse{
		Status:  status,
		Name:    name,
		Payload: bodyBytes,
	}
	return res, nil
}

func (*server) WebCrawlerBatch(stream pb.WebCrawlerService_WebCrawlerBatchServer) error {
	fmt.Printf("WebCrawlerBatch function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		cnResponse := make(chan *pb.WebCrawlerBatchResponse, 5)
		lenBatchReq := len(req.GetRequests())

		// processing batch requests in parallel
		for _, r := range req.GetRequests() {
			go func(r *pb.WebCrawlerBatchRequest_Request) {
				url := r.GetUrl()
				fmt.Println("Got URL to request: " + url + "! ")
				resp, st := Crawler(url)

				cnResponse <- &pb.WebCrawlerBatchResponse{
					Payload: []byte(resp),
					Name:    url,
					Status:  st,
				}
			}(r)
		}

		// sending results assync
		for i := 0; i < lenBatchReq; i++ {
			r := <-cnResponse
			sendErr := stream.Send(r)
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
			}
		}
	}
}

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
	pb.RegisterWebCrawlerServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
