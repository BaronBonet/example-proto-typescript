package main

import (
	"context"
	"fmt"
	examplev1 "github.com/BaronBonet/dummy-backend/generated/go/example/v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
)

type exampleServiceServer struct {
	examplev1.UnimplementedExampleServiceServer
}

func (s *exampleServiceServer) GetExampleMessage(ctx context.Context, req *examplev1.GetExampleMessageRequest) (*examplev1.GetExampleMessageResponse, error) {
	log.Println("I got a message request: request")
	ranNumber := req.GetRanNumber()
	concatString := fmt.Sprintf("%d_example", ranNumber)

	exampleMessage := &examplev1.ExampleMessage{
		Number:       ranNumber,
		ConcatString: concatString,
	}

	res := &examplev1.GetExampleMessageResponse{
		ExampleMessage: exampleMessage,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	examplev1.RegisterExampleServiceServer(grpcServer, &exampleServiceServer{})
	log.Println("gRPC server started on port 8080")
	// Enable gRPC reflection
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
