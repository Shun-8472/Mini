package demo

import (
	"context"

	pb "demo/external/protos/demo/v1"
	proc "demo/internal/processor/version/procedure"
)

type impl struct {
	proc proc.Procedure
}

func (im *impl) WhatIsDemoInfo(ctx context.Context, request *pb.WhatIsDemoInfoRequest) (*pb.WhatIsDemoInfoResponse, error) {
	demoInfo, err := im.proc.GetDemoInfo()
	if err != nil {
		return nil, err
	}
	resultDemoInfo := pb.WhatIsDemoInfoResponse{
		Version:          demoInfo.Version,
		LatestUpdateTime: demoInfo.LastUpdate,
	}

	return &resultDemoInfo, nil
}

func ProvideReceiver(proc proc.Procedure) Receiver {
	return &impl{
		proc: proc,
	}
}
