package crawler

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/mtulio/grpc-services/src/protobuf/pbcrawler"
	"google.golang.org/grpc"
)

type service struct{}

// Register the service on the gRPC server
func Register(gs *grpc.Server) error {
	pb.RegisterWebCrawlerServiceServer(gs, &service{})

	return nil
}

// WebCrawler is the gRPC server interface implementation.
func (*service) WebCrawler(ctx context.Context, req *pb.WebCrawlerRequest) (*pb.WebCrawlerResponse, error) {
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
			bodyString = fmt.Sprintf("error, statusCode: %d", response.StatusCode)
		} else {
			bodyString = string(bodyBytes)
			status = true
		}
	} else {
		bodyString = fmt.Sprintf("error, statusCode: %d", response.StatusCode)
	}

	fmt.Println("Body length: ", len(bodyString))
	res := &pb.WebCrawlerResponse{
		Status:  status,
		Name:    name,
		Payload: bodyBytes,
	}
	return res, nil
}

// WebCrawlerBatch is the gRPC server interface implementation.
func (*service) WebCrawlerBatch(stream pb.WebCrawlerService_WebCrawlerBatchServer) error {
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
		if len(req.Requests) == 0 {
			return nil
		}
		fmt.Println(len(req.Requests))

		cnResponse := make(chan *pb.WebCrawlerBatchResponse, 5)
		lenBatchReq := len(req.GetRequests())

		// processing batch requests in parallel
		for _, r := range req.GetRequests() {
			go func(r *pb.WebCrawlerBatchRequest_Request) {
				url := r.GetUrl()
				fmt.Println("Got URL to request: " + url + "! ")
				resp, st := crawler(url)

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

func crawler(url string) (string, bool) {
	status := false
	bodyString := ""
	response, err := http.Get(url)
	var bodyBytes []byte

	if err != nil {
		return fmt.Sprintf("error, statusCode: %s", err), status
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			bodyString = fmt.Sprintf("error, statusCode: %d", response.StatusCode)
		} else {
			bodyString = string(bodyBytes)
			status = true
		}
	} else {
		bodyString = fmt.Sprintf("error, statusCode: %d", response.StatusCode)
	}

	return bodyString, status
}
