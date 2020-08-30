package main

import (
    "google.golang.org/grpc"
    "github.com/pkg/errors"
    "log"
    "net"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterBookInfoServer(s, &server{})

    log.Printf("Starting gRPC listener on port " + port)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}