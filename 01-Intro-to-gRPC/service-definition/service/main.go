package main

import (
	"fmt"
	"net"
	"os"

	pb "github.com/petrostrak/gRPC-comprehensive/01-Intro-to-gRPC/service-definition/service/commerce"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := grpc.Server{}
	pb.RegisterProductInfoServer(s, &server{})

	err = s.Serve(l)
	if err == nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
