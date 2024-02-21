package procedure

import (
	"github.com/google/wire"

	"demo/internal/processor/version/implement"
)

var ProcSet = wire.NewSet(implement.NewDemoProcedure)
