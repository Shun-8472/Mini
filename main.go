package main

import (
	"mini/config"
	"mini/inject"
)

func main() {
	//init config
	config.InitConfigs()

	//build inject
	injects, err := inject.BuildInjector()
	if err != nil {
		panic("failed to inject")
	}

	//Connect Cache
	//injects.Redis.ConnectCache()

	//Connect Database
	//injects.Database.ConnectDatabase()

	//Connect LLM
	injects.LLM.ConnectLLM()

	//Start GRPC Server
	injects.GrpcServer.StartGRPCServer(injects.VersionReceiver, injects.ChatReceiver)
}
