package engine

import "demo/external/receiver/demo"

type Engine interface {
	StartGRPCServer(sve demo.Receiver)
}
