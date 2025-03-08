//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"

	"mini/inject/applied"
	"mini/inject/external"
	"mini/inject/internal/procedure"
)

func BuildInjector() (*Injects, error) {
	wire.Build(
		procedure.ProcSet,

		external.MiniSet,

		applied.InitGrpcServer,
		applied.InitDatabaseConnection,
		applied.InitCacheConnection,
		applied.InitLLMConnection,

		InjectsSet,
	)
	return new(Injects), nil
}
