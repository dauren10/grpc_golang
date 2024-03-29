package main

import (
	"fmt"
	"log"
	"net"

	"context"

	"github.com/dauren10/demo-grpc/myserver"
	"google.golang.org/grpc"
)

type MyServiceServer struct {
	myserver.UnimplementedMyServiceServer
}

func (s MyServiceServer) Create(ctx context.Context, req *myserver.CreateRequest) (*myserver.CreateResponse, error) {

	response := &myserver.CreateResponse{
		Number: []byte("hello"),
		Price:  []byte("hi"),
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	} else {
		fmt.Println("ff")
	}
	serverRegister := grpc.NewServer()
	service := &MyServiceServer{}
	myserver.RegisterMyServiceServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
}
