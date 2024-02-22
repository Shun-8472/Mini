package main

import (
	"demo/config"
	"demo/inject"
)

func main() {
	//init config
	config.InitConfigs()

	//build inject
	injects, err := inject.BuildInjector()
	if err != nil {
		panic("failed to inject")
	}

	//Connect Database
	injects.Database.ConnectDatabase()
	//Start GRPC Server
	injects.GrpcServer.StartGRPCServer(injects.Receiver)
}
