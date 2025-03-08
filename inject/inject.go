package inject

import (
	"github.com/google/wire"

	chatReciver "mini/external/receiver/chat"
	versionReciver "mini/external/receiver/version"
	"mini/internal/applied/cache"
	"mini/internal/applied/database"
	"mini/internal/applied/engine"
	"mini/internal/applied/llm"
)

var InjectsSet = wire.NewSet(wire.Struct(new(Injects), "*"))

type Injects struct {
	VersionReceiver versionReciver.Receiver
	GrpcServer      engine.Engine
	Database        database.Database
	Redis           cache.Cache

	//LLM Receiver
	ChatReceiver chatReciver.Receiver
	LLM          llm.LLM
}
