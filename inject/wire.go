//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"

	"demo/inject/applied"
	"demo/inject/external"
	"demo/inject/internal/procedure"
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
