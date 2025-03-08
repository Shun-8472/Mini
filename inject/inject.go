package inject

import (
	"github.com/google/wire"
	"mini/external/receiver/demo"

	"mini/external/receiver/chat"
	"mini/internal/applied/cache"
	"mini/internal/applied/database"
	"mini/internal/applied/engine"
	"mini/internal/applied/llm"
)

var InjectsSet = wire.NewSet(wire.Struct(new(Injects), "*"))

type Injects struct {
	Receiver   demo.Receiver
	GrpcServer engine.Engine
	Database   database.Database
	Redis      cache.Cache

	//LLM Receiver
	ChatReceiver chat.Receiver
	LLM          llm.LLM
}
