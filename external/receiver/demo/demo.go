package demo

import pb "mini/external/protos/demo/v1"

type Receiver interface {
	pb.DemoServiceServer
}
