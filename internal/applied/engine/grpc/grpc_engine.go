package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"demo/config"
	pb "demo/external/protos/demo/v1"
	"demo/external/receiver/demo"
	"demo/internal/applied/engine"
)

type Engine struct {
}

func NewEngine() engine.Engine {
	return &Engine{}
}

func (e Engine) StartGRPCServer(sve demo.Receiver) {
	address := config.GetServerAddress()
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic("failed to listen address")
	}

	s := grpc.NewServer()

	pb.RegisterDemoServiceServer(s, sve)
	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		panic("failed to start server")
	}
}
