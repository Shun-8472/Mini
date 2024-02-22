package external

import (
	"github.com/google/wire"

	"demo/external/receiver/demo"
)

var DemoSet = wire.NewSet(demo.ProvideReceiver)
