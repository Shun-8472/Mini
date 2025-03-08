package version

import (
	version2 "mini/internal/processor/version/procedure/version"
	"mini/internal/processor/version/substance/version"
)

type Procedure struct {
}

func NewVersionProcedure() version2.Procedure {
	return &Procedure{}
}

func (p Procedure) GetVersionInfo() (version.DemoInfo, error) {
	return version.DemoInfo{
		Version:    "v1",
		LastUpdate: "2024-02-22",
	}, nil
}
