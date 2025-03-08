package version

import (
	"context"

	pb "mini/external/protos/version/v1"
	proc "mini/internal/processor/version/procedure/version"
)

type impl struct {
	proc proc.Procedure
}

func (im *impl) WhatIsVersionInfo(ctx context.Context, request *pb.WhatIsVersionInfoRequest) (*pb.WhatIsVersionInfoResponse, error) {
	demoInfo, err := im.proc.GetVersionInfo()
	if err != nil {
		return nil, err
	}
	resultDemoInfo := pb.WhatIsVersionInfoResponse{
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
