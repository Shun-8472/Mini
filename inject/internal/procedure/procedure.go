package procedure

import (
	"github.com/google/wire"

	"mini/internal/processor/version/implement/chat"
	"mini/internal/processor/version/implement/version"
)

var ProcSet = wire.NewSet(version.NewDemoProcedure, chat.NewChatProcedure)
