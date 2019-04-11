package main

import (
	"context"
	"log"
	"net"

	"github.com/payfazz/test-grpc/proto"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":16969")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &Server{})
	reflection.Register(srv)

	log.Println("Starting to serve the grpc server")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

type Server struct{}

func (s *Server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b, c := request.GetA(), request.GetB(), request.GetC()

	result := a + b + c

	log.Println("grpc result for add : ", result)
	return &proto.Response{Result: result}, nil
}

func (s *Server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b, c := request.GetA(), request.GetB(), request.GetC()

	result := a * b * c

	log.Println("grpc result for multiply 	: ", result)
	return &proto.Response{Result: result}, nil
}

func (s *Server) Minus(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b, c := request.GetA(), request.GetB(), request.GetC()

	result := a - b - c

	log.Println("grpc result for minus	: ", result)
	return &proto.Response{Result: result}, nil
}
