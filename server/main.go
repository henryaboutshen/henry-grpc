package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "henry_grpc/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type serverHello struct {
	pb.UnimplementedHelloServer
}

type serverHi struct {
	pb.UnimplementedHiServer
}

func InternalFunc() {
	log.Printf("Internal")
	return
}

func (s *serverHello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	InternalFunc()
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *serverHi) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiReply, error) {
	log.Printf("Received: %v", in.GetName())
	InternalFunc()
	return &pb.HiReply{Message: "Hi " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s1 := grpc.NewServer()
	pb.RegisterHelloServer(s1, &serverHello{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s1.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	s2 := grpc.NewServer()
	pb.RegisterHiServer(s2, &serverHi{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s2.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
