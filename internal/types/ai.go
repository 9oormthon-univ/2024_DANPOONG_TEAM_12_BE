package types

type AIService interface {
	GenerateTrainingData() error
	RequestFineTuning() error
	InjectInfoService(service RegionsService)
	RecommendCourses(req *RecommendCourseRequest) (*AnswerResponse, error)
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatFormat struct {
	Messages []*ChatMessage `json:"messages"`
}

const TrainingDataPath = "/app/internal/data/output.jsonl"

type AnswerResponse struct {
	Title string `json:"title"`
}

type RecommendCourseRequest struct {
	Interests []string `json:"interests"`
	AreaCode  AreaCode `json:"areaCode"`
}
