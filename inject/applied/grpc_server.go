package applied

import (
	"mini/internal/applied/engine"
	"mini/internal/applied/engine/grpc"
)

func InitGrpcServer() engine.Engine {
	return grpc.NewEngine()
}
