package types

type AIService interface {
	InjectInfoService(service RegionsService)
	// DefineFunctions() []openai.FunctionDefinition
	RecommendCourses(req *RecommendCourseRequest) ([]*TravelRecommendation, error)
	GetTourRecommendations(region string, interests []string) (string, error)
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
	Answer string `json:"content"`
}
type TravelRecommendation struct {
	Title       string `json:"title"`       // 장소 이름
	Description string `json:"description"` // 장소 설명
	Address     string `json:"address"`     // 장소 주소
	StartTime   string `json:"start_time"`  // 시작 시간 (예: "10:00")
	EndTime     string `json:"end_time"`
	Type        string `json:"type"`
}

type RecommendCourseRequest struct {
	AreaCode  string   `json:"area_code"`  // 지역 코드
	Interests []string `json:"interests"`  // 관심사
	StartTime string   `json:"start_time"` // 여행 시작 시간 (예: "10:00")
	EndTime   string   `json:"end_time"`   // 여행 종료 시간 (예: "18:00")
}
