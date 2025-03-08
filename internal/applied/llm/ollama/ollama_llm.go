package ollama

import (
	"github.com/tmc/langchaingo/llms/ollama"

	"mini/config"
)

var (
	LLMConnection = new(ollama.LLM)
)

type LLM struct {
}

func NewConnection() *LLM {
	return &LLM{}
}

func (l *LLM) ConnectLLM() {
	client, err := ollama.New(ollama.WithModel(config.C.LLM.Ollamamodel))

	if err != nil {
		panic(err)
	}
	LLMConnection = client
}
