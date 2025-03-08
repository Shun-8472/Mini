package engine

import (
	"mini/external/receiver/chat"
	"mini/external/receiver/demo"
)

type Engine interface {
	StartGRPCServer(sve demo.Receiver, chatSve chat.Receiver)
}
