package interfaces

type LLMService interface {
	GenerateComment(newsContent string, prompt string) (string, error)
}
