package types

type AIService interface {
	GenerateTrainingData() error
	InjectInfoService(service RegionsService)
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatFormat struct {
	Messages []*ChatMessage `json:"messages"`
}
