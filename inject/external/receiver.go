package external

import (
	"github.com/google/wire"

	"mini/external/receiver/chat"
	"mini/external/receiver/demo"
)

var MiniSet = wire.NewSet(demo.ProvideReceiver, chat.ProvideReceiver)
