package version

import pb "mini/external/protos/version/v1"

type Receiver interface {
	pb.VersionServiceServer
}
