package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"service/internal/generated/logprinter"

	"google.golang.org/grpc"
)

type server struct {
	logprinter.UnimplementedPrinterServiceServer
}

func (s *server) PrintLog(ctx context.Context, req *logprinter.LogRequest) (*logprinter.LogResponse, error) {
	log.Printf("received request: %v", req.GetMsg())
	return &logprinter.LogResponse{
		Size: int32(len(req.GetMsg())),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 80))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// You are likely going to bind your server implementation here
	logprinter.RegisterPrinterServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
