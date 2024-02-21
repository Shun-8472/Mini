package implement

import (
	"demo/internal/processor/version/procedure"
	"demo/internal/processor/version/substance"
)

type Procedure struct {
}

func NewDemoProcedure() procedure.Procedure {
	return &Procedure{}
}

func (p Procedure) GetDemoInfo() (substance.DemoInfo, error) {
	return substance.DemoInfo{
		Version:    "v1",
		LastUpdate: "2024-02-22",
	}, nil
}
