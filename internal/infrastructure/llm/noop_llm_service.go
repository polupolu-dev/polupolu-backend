package llm

import (
	"context"
)

type NoopLLMService struct{}

func (n *NoopLLMService) GenerateComment(ctx context.Context, text, prompt string) (string, error) {
	return "", nil
}
