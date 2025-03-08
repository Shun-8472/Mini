package chat

import pb "mini/external/protos/chat/v1"

type Receiver interface {
	pb.ChatServiceServer
}
