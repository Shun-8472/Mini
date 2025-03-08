package demo

import (
	"context"

	pb "mini/external/protos/demo/v1"
	proc "mini/internal/processor/version/procedure/version"
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
