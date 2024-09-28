package main

import (
    "context"
    "log"
    "net"

    pb "github.com/pardocch/grpc-server/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    log.Printf("Received: %v", req.GetName())
    return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterGreeterServer(grpcServer, &server{})

    log.Printf("Server is listening on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

