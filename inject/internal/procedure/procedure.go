package procedure

import (
	"github.com/google/wire"

	chatProcedure "mini/internal/processor/version/implement/chat"
	versionProcedure "mini/internal/processor/version/implement/version"
)

var ProcSet = wire.NewSet(versionProcedure.NewVersionProcedure, chatProcedure.NewChatProcedure)
