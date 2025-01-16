package interfaces

import (
	"context"
)

type LLMService interface {
	GenerateComment(ctx context.Context, newsContent string, prompt string) (string, error)
}
