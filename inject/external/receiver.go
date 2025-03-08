package external

import (
	"github.com/google/wire"

	chatReciver "mini/external/receiver/chat"
	versionReciver "mini/external/receiver/version"
)

var MiniSet = wire.NewSet(versionReciver.ProvideReceiver, chatReciver.ProvideReceiver)
