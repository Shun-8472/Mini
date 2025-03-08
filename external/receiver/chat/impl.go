package chat

import (
	"context"

	"github.com/tmc/langchaingo/llms"

	pb "mini/external/protos/chat/v1"
	"mini/internal/applied/llm/ollama"
	proc "mini/internal/processor/version/procedure/chat"
)

type impl struct {
	proc proc.Procedure
}

func (i impl) GenerateMessage(ctx context.Context, request *pb.ChatRequest) (*pb.ChatResponse, error) {
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, request.UserInput),
	}

	result, err := ollama.LLMConnection.GenerateContent(ctx, content)
	if err != nil {
		return nil, err
	}

	var responseText string
	if len(result.Choices) > 0 {
		responseText = result.Choices[0].Content
	} else {
		responseText = "No response generated."
	}

	return &pb.ChatResponse{Response: responseText}, nil
}

func (i impl) StreamMessage(request *pb.ChatRequest, stream pb.ChatService_StreamMessageServer) error {
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, request.UserInput),
	}

	_, err := ollama.LLMConnection.GenerateContent(context.Background(), content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		if err := stream.Send(&pb.ChatResponse{Response: string(chunk)}); err != nil {
			return err
		}
		return nil
	}))
	return err
}

func ProvideReceiver(proc proc.Procedure) Receiver {
	return &impl{
		proc: proc,
	}
}
