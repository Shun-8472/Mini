package grpc

import (
	"mini/external/receiver/chat"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"mini/config"
	chatPb "mini/external/protos/chat/v1"
	pb "mini/external/protos/demo/v1"
	"mini/external/receiver/demo"
	"mini/internal/applied/engine"
)

type Engine struct {
}

func NewEngine() engine.Engine {
	return &Engine{}
}

func (e Engine) StartGRPCServer(sve demo.Receiver, chatSve chat.Receiver) {
	address := config.GetServerAddress()
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic("failed to listen address")
	}

	s := grpc.NewServer()

	pb.RegisterDemoServiceServer(s, sve)
	chatPb.RegisterChatServiceServer(s, chatSve)

	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		panic("failed to start server")
	}
}
