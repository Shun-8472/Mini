//go:build wireinject
// +build wireinject

package inject

import (
	applied2 "demo/inject/applied"
	"demo/inject/external"
	"demo/inject/internal/procedure"
	"github.com/google/wire"
)

func BuildInjector() (*Injects, error) {
	wire.Build(
		procedure.ProcSet,

		external.DemoSet,

		applied2.InitGrpcServer,
		applied2.InitDatabaseConnect,

		InjectsSet,
	)
	return new(Injects), nil
}
