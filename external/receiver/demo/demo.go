package demo

import pb "demo/external/protos/demo/v1"

type Receiver interface {
	pb.DemoServiceServer
}
