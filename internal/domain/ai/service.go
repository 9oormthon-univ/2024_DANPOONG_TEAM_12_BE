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
	types.CarpoolsService
	types.MatchingService
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

// id, details, categories + location(사용자 위치)
// 컨트롤러에서 사용자 지역 위치 반환
func (service *aiService) RecommendMatchingPost(page int, pageSize int, location string, interests []string) ([]*types.MatchingDetailForAI, error) {
	// 1. 매칭 게시글 조회
	// posts, err := service.MatchingService.GetPostsForAI(page, pageSize)
	posts, err := service.GetExampleMatchingPosts()

	if err != nil {
		return nil, fmt.Errorf("게시글 조회 오류: %w", err)
	}

	// 데이터가 없는 경우 처리
	if len(posts) == 0 {
		return nil, fmt.Errorf("게시글이 없음")
	}

	// 2. 프롬프트 생성
	prompt := fmt.Sprintf(`
나의 위치는 %s이고, 관심사는 %s이야. 아래의 매칭 게시글들을 참고해서 나에게 가장 적합한 매칭 게시글 1개를 추천해줘. 추천된 게시글의 ID, 상세 내용, 카테고리를 JSON 배열 형식으로 반환해줘.

매칭 게시글들:
`, location, strings.Join(interests, ", "))

	// 각 게시글의 상세 정보 추가
	for _, post := range posts {
		categoriesStr := strings.Join(post.Categories, ", ")
		prompt += fmt.Sprintf(`
---
게시글 ID: %s
상세 내용: %s
카테고리: %s
`, post.MatchingID, post.Details, categoriesStr)
	}

	// 프롬프트 마무리
	prompt += `
---
응답 형식:
[
    {
        "matching_id": "숫자",
        "details": "문자열",
        "categories": ["카테고리1", "카테고리2"]
    },
    ...
]
`

	// 3. OpenAI 메시지 구성
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "사용자의 근처 동네와 관심사를 기반으로 매칭 게시글의 카테고리와 내용을 사용자에게 맞게 추천해주는 어시스턴트입니다. 최종적으로 게시글 ID, 상세 내용, 카테고리를 JSON 배열 형식으로 제공합니다.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	// 4. OpenAI API 요청 구성
	req := openai.ChatCompletionRequest{
		Model:       openai.GPT4Turbo, // 사용하려는 모델로 변경 가능
		Messages:    messages,
		MaxTokens:   500, // 응답의 길이에 따라 조정
		Temperature: 0.5, // 창의성 조절
	}

	// 5. OpenAI API 호출
	resp, err := service.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Printf("Error calling OpenAI API: %v", err)
		return nil, fmt.Errorf("OpenAI API 호출 오류: %w", err)
	}

	// 6. 응답 검증
	if len(resp.Choices) == 0 {
		log.Println("No recommendations returned from OpenAI.")
		return nil, fmt.Errorf("OpenAI로부터 추천을 받을 수 없습니다.")
	}

	// 7. 응답에서 추천된 게시글 정보 추출
	recommendedContent := strings.TrimSpace(resp.Choices[0].Message.Content)

	// 응답 로깅 (디버깅 용도)
	log.Printf("AI 응답 내용: %s", recommendedContent)

	// 8. JSON 파싱 (추천된 게시글 정보가 JSON 배열 형식이라고 가정)
	var recommendedPosts []*types.MatchingDetailForAI
	err = json.Unmarshal([]byte(recommendedContent), &recommendedPosts)
	if err != nil {
		log.Printf("Error parsing AI response as JSON: %v", err)
		return nil, fmt.Errorf("AI 응답 파싱 오류: %w", err)
	}

	// 9. 추천 게시글이 없는 경우 처리
	if len(recommendedPosts) == 0 {
		log.Println("AI did not return any recommended posts.")
		return nil, fmt.Errorf("추천할 게시글이 없습니다.")
	}

	// 10. 추천 게시글 반환
	return recommendedPosts, nil
}

func (a *aiService) InjectRegionService(service types.RegionsService) {
	a.RegionsService = service
}
func (a *aiService) InjectCarpoolsService(service types.CarpoolsService) {
	a.CarpoolsService = service
}
func (a *aiService) InjectMatchingService(service types.MatchingService) {
	a.MatchingService = service
}
