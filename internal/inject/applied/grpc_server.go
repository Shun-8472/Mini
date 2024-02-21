package applied

import (
	"demo/internal/applied/engine"
	"demo/internal/applied/engine/grpc"
)

func InitGrpcServer() engine.Engine {
	return grpc.NewEngine()
}
