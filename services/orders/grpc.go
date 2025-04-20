package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr); if err != nil {
		log.Fatalf("failed to listen: %v", err);
	}
	gRPCServer := grpc.NewServer()
	// register gRPC out services here

	log.Println("Starting gRPC Server on", s.addr);

	return gRPCServer.Serve(lis);
}