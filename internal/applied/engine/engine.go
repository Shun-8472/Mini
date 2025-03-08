package engine

import (
	"mini/external/receiver/chat"
	"mini/external/receiver/version"
)

type Engine interface {
	StartGRPCServer(sve version.Receiver, chatSve chat.Receiver)
}
