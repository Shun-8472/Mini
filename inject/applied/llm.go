package applied

import (
	"mini/internal/applied/llm"
	"mini/internal/applied/llm/ollama"
)

func InitLLMConnection() llm.LLM {
	return ollama.NewConnection()
}
