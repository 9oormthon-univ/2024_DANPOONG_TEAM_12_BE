package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type aiService struct {
	aiRepository *AIRepository
	types.RegionsService
	client *openai.Client
}

// TravelRecommendation 구조체에 Type 필드 추가

func SetAIService(repository *AIRepository) types.AIService {
	a := &aiService{
		aiRepository: repository,
	}
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		log.Println("key : ", key)
		panic("openai api key가 없음")
	}

	a.client = openai.NewClient(key)
	return a
}

func (a *aiService) DefineFunctions() []openai.FunctionDefinition {
	return []openai.FunctionDefinition{
		{
			Name:        "get_tour_recommendations",
			Description: "사용자의 지역, 관심사, 시간대에 맞춰 Tour API 카테고리로 매핑한 데이터를 반환합니다.",
			Parameters: &jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"region": {
						Type:        jsonschema.String,
						Description: "사용자가 방문하고 싶은 지역 (예: 전라북도)",
					},
					"interests": {
						Type: jsonschema.String,
						Description: fmt.Sprintf(
							"사용자가 자유롭게 입력한 관심사 목록입니다. 다음과 같이 입력해 관심사 마다 공백으로 구분할거야 -> '문화시설 관광지'. AI는 이를 다음의 Tour API 카테고리로 매핑해야 합니다: [%s].",
							strings.Join(types.ContentTypeNames, ", "),
						),
					},
				},
				Required: []string{"region", "interests"},
			},
		},
	}
}

func (a *aiService) RecommendCourses(req *types.RecommendCourseReq) ([]*types.CourseRecommendationAIRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 함수 정의 가져오기
	functions := a.DefineFunctions()

	// AI 요청 메시지
	messages := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf(
				`당신은 여행 코스를 추천하는 어시스턴트입니다.
사용자가 입력한 관심사를 다음의 Tour API 카테고리로 값들 중 최대 3개 이하로 바꿔서 함수 매개변수로 사용해야 합니다: [%s].
제공된 데이터를 기반으로 시간 순서대로 여행 코스를 생성하고, 각 장소의 시작 시간과 끝 시간을 자동으로 배정하세요.
응답은 배열 형태의 JSON으로만 반환해야 하며, 각 장소는 'title', 'description', 'address', 'start_time', 'end_time', 'type', 'content_id'을 포함해야 합니다.
응답은 반드시 배열 형태의 JSON으로만 시작하고 끝나야 합니다. 추가적인 텍스트나 설명을 포함하지 마세요.
예를 들어:
[
  {
    "title": "장소 이름",
    "description": "장소 설명",
    "address": "주소",
    "start_time": "09:00",
    "end_time": "10:30",
    "type": "관광지"
	"content_id": 331345",
  }
]`,
				strings.Join(types.ContentTypeNames, ", "),
			),
		},
		{
			Role: openai.ChatMessageRoleUser,
			Content: fmt.Sprintf(
				"지역: %s, 관심사: %s, 여행 시간: %s부터 %s까지. 관심사를 위의 카테고리로 바꿔서 적합한 장소를 추천해주세요.",
				a.RegionsService.GetAreaNameByCode(types.AreaCode(req.AreaCode)),
				req.Categories,
				req.StartTime,
				req.EndTime,
			),
		},
	}

	// AI 요청 및 재시도 로직 추가
	const maxRetries = 3
	var res openai.ChatCompletionResponse
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		res, err = a.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:        openai.GPT4Turbo,
			Messages:     messages,
			Functions:    functions,
			FunctionCall: "auto",
		})
		if err == nil {
			break
		}
		log.Printf("AI 요청 시도 %d/%d 실패: %v", attempt, maxRetries, err)
		time.Sleep(2 * time.Second) // 재시도 전 대기
	}
	if err != nil {
		return nil, fmt.Errorf("AI 요청 오류: %w", err)
	}

	// 함수 호출 처리
	message := res.Choices[0].Message

	if message.FunctionCall != nil {
		var args struct {
			Region    string `json:"region"`
			Interests string `json:"interests"`
		}
		if err := json.Unmarshal([]byte(message.FunctionCall.Arguments), &args); err != nil {
			return nil, fmt.Errorf("함수 인자 파싱 오류: %w", err)
		}

		// 관심사 매핑 로그
		log.Printf("%s -> %s", req.Categories, args.Interests)

		// 관심사 매핑: 문자열을 공백으로 분리하여 배열로 변환
		interests := strings.Fields(args.Interests)

		// Tour API 데이터 호출
		recommendations, err := a.GetTourRecommendations(args.Region, interests)
		if err != nil {
			return nil, fmt.Errorf("Tour API 호출 오류: %w", err)
		}

		// Tour API 결과를 AI에게 다시 전달
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleFunction,
			Name:    "get_tour_recommendations",
			Content: recommendations,
		})

		// AI가 최종 코스를 생성하도록 요청
		finalRes, err := a.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:    openai.GPT4Turbo, // 올바른 모델 이름 사용
			Messages: messages,
		})
		if err != nil {
			return nil, fmt.Errorf("최종 AI 응답 오류: %w", err)
		}

		// 최종 응답 처리
		var finalRecommendations []*types.CourseRecommendationAIRes

		// JSON만 추출하기 위해 정규 표현식 사용
		re := regexp.MustCompile(`(?s)\[.*\]`)
		jsonMatch := re.FindString(finalRes.Choices[0].Message.Content)
		if jsonMatch == "" {
			log.Printf("JSON 형식이 포함되지 않은 응답: %s", finalRes.Choices[0].Message.Content)
			return nil, fmt.Errorf("응답에 JSON 형식이 포함되지 않았습니다")
		}

		if err := json.Unmarshal([]byte(jsonMatch), &finalRecommendations); err != nil {
			log.Printf("JSON 파싱 오류 데이터: %s", jsonMatch)
			return nil, fmt.Errorf("최종 JSON 파싱 오류: %w", err)
		}

		return finalRecommendations, nil
	}

	return nil, fmt.Errorf("AI가 함수 호출을 수행하지 않았습니다")
}

// Tour API 데이터를 반환하는 함수(AI가 매개변수 입력)
func (a *aiService) GetTourRecommendations(region string, interests []string) (string, error) {
	areaCode := a.RegionsService.GetAreaCodeByName(region)
	if areaCode == "" {
		return "", fmt.Errorf("지역 코드 매핑 실패")
	}

	var allRecommendations []*types.CourseRecommendationAIRes
	for _, interest := range interests {
		log.Printf("매핑된 관심사: %s", interest)
		contentTypeID := a.RegionsService.GetContentTypeCodeByName(interest)
		if contentTypeID == "" {
			log.Printf("관심사 '%s'에 대한 콘텐츠 타입 ID 매핑 실패", interest)
			continue
		}

		// 관심사별 데이터 호출
		areaList, err := a.RegionsService.GetAreaBasedList(areaCode, contentTypeID)
		if err != nil {
			log.Printf("'%s'에 대한 데이터 가져오기 오류: %v", interest, err)
			continue
		}

		// 상세 데이터 가져오기
		for _, area := range areaList {
			detail, err := a.RegionsService.GetDetailCommon(area.ContentID)
			if err != nil {
				log.Printf("상세 정보 가져오기 실패: %s", err)
				continue
			}

			// 데이터 추가
			allRecommendations = append(allRecommendations, &types.CourseRecommendationAIRes{
				Title:       detail.Title,
				Description: detail.Overview,
				Address:     detail.Addr1,
				StartTime:   "", // 시간은 AI에서 처리
				EndTime:     "",
				Type:        "", // 타입은 AI에서 처리
				ContentID:   detail.ContentID,
			})
		}
	}

	// JSON 형태로 반환
	recommendationsJSON, err := json.Marshal(allRecommendations)
	if err != nil {
		return "", fmt.Errorf("JSON 변환 오류: %w", err)
	}

	return string(recommendationsJSON), nil
}

func (a *aiService) InjectRegionService(service types.RegionsService) {
	a.RegionsService = service
}
