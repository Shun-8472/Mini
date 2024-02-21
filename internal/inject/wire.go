//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"

	"demo/internal/inject/applied"
	"demo/internal/inject/external"
	"demo/internal/inject/procedure"
)

func BuildInjector() (*Injects, error) {
	wire.Build(
		procedure.ProcSet,

		external.DemoSet,

		applied.InitGrpcServer,
		applied.InitDatabaseConnect,

		InjectsSet,
	)
	return new(Injects), nil
}
