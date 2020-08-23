package main

import (
	"context"
	"github.com/EDEN45/Blog-Application/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type authServer struct {
}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	return &proto.AuthResponse{}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal("Error create listener: ", err.Error())
	}

	server.Serve(listener)
}
