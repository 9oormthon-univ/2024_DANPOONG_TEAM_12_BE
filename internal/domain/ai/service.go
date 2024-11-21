package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type aiService struct {
	aiRepository *AIRepository
	types.RegionsService
	client        *openai.Client
	fineTuneModel string
}

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

func (a *aiService) RequestFineTuning() error {
	f, err := a.client.CreateFile(context.TODO(), openai.FileRequest{
		FilePath: types.TrainingDataPath,
		Purpose:  "fine-tune",
	})
	if err != nil {
		return fmt.Errorf("Create file error: %v\n", err)
	}

	_, err = os.Stat(types.TrainingDataPath)
	if os.IsNotExist(err) {
		fmt.Println("output.jsonl 파일이 현재 디렉토리에 존재하지 않습니다.")
	}

	if err != nil {
		return fmt.Errorf("Upload JSONL file error: %v\n", err)
	}

	fineTuningJob, err := a.client.CreateFineTuningJob(context.TODO(), openai.FineTuningJobRequest{
		TrainingFile: f.ID,
		Model:        openai.GPT4oMini20240718,
	})
	if err != nil {
		return fmt.Errorf("Creating new fine-tune model error: %v\n", err)
	}
	for {
		fineTuningJob, err = a.client.RetrieveFineTuningJob(context.TODO(), fineTuningJob.ID)
		if err != nil {
			log.Printf("Getting fine-tune model error: %v\n", err)
			return err
		}

		if fineTuningJob.Status == "succeeded" {
			fmt.Println("Fine-tuned model created successfully:", fineTuningJob.FineTunedModel)
			a.fineTuneModel = fineTuningJob.FineTunedModel
			return nil
		}

		fmt.Println("Waiting for fine-tuning job to complete...")
		time.Sleep(10 * time.Second)
	}
}

// Start of Selection

func (a *aiService) RecommendCourses(req *types.RecommendCourseRequest) (*types.AnswerResponse, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	var result *types.AnswerResponse
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		return nil, fmt.Errorf("Schema generation error %w\n", err)
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "사용자 관심사에 맞는 지역 여행 코스를 추천해주는 서비스입니다",
		},
		{
			Role: openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("나의 관심사는 %s이고, 위치는 %s이야 조건에 맞게 지역 여행 코스를 추천해줘.",
				req.Interests,
				a.RegionsService.GetAreaNameByCode(req.AreaCode),
			),
		},
	}

	res, err := a.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    "ft:gpt-4o-2024-08-06:personal::ASEB50GB",
		Messages: messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{

			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "RegionalTravelInfo",
				Schema: schema,
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Create completion error %w\n", err)
	}
	if err := schema.Unmarshal(res.Choices[0].Message.Content, &result); err != nil {
		return nil, fmt.Errorf("Unmarshal error %w\n", err)
	}

	return result, nil
}

func (a *aiService) GenerateTrainingData() error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "/internal/data/output.jsonl")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// 관심사 하나당 관심사에 맞는 지역 정보 2개씩 생성
	var chatFormatList []*types.ChatFormat
	for _, interest := range types.ContentTypes {
		for _, areaCode := range types.AreaCodes {
			log.Println("데이터 생성 중.....")
			areaList, err := a.RegionsService.GetareaBasedList(areaCode, interest)
			if err != nil {
				return err
			}

			for _, area := range areaList {
				detail, err := a.RegionsService.GetDetailCommon(area.ContentID)
				if err != nil {
					return err
				}
				if detail.Overview == "" {
					log.Println("여행지 아님 : ", detail.Title, detail.ContentID)
					continue
				}
				// 숙박 정보, 관심사와 위치에 따른 정보 2~3개
				log.Printf("areaCode : %s", detail.AreaCode)
				log.Printf("contentid %s : ", detail.ContentID)
				log.Printf("contenttypeid : %s", detail.ContentTypeID)
				msg := &types.ChatFormat{
					Messages: []*types.ChatMessage{
						{
							Role:    openai.ChatMessageRoleSystem,
							Content: "사용자 관심사에 맞게 Tour API 내의 관광지, 문화시설, 축제공연행사, 여행코스, 레포츠, 숙박, 쇼핑, 음식점을 사용자에게 제공하는 서비스입니다. 해당 API 외의 정보는 제공하지 않으며, 학습한 데이터와 Tour API 내의 데이터만으로 추천을 진행합니다.",
						},
						{
							Role: openai.ChatMessageRoleUser,
							Content: fmt.Sprintf("내가 가고 싶은 지역은 %s이고, 관심사는 %s이야. 조건에 맞게 지역 여행지(관광지, 문화시설, 축제공연행사, 여행코스, 레포츠, 숙박, 쇼핑, 음식점)를 추천해줘",
								a.RegionsService.GetAreaNameByCode(types.AreaCode(detail.AreaCode)),
								a.RegionsService.GetContentTypeNameByCode(types.ContentType(detail.ContentTypeID))),
						},
						{
							Role:    openai.ChatMessageRoleAssistant,
							Content: fmt.Sprintf("타이틀: %s, 간단한 소개: %s", detail.Title, detail.Overview),
						},
					},
				}
				chatFormatList = append(chatFormatList, msg)

				data, err := json.Marshal(msg)
				if err != nil {
					return fmt.Errorf("json 마샬 오류 : %w\n", err)
				}
				if _, err = f.WriteString(fmt.Sprintf("%s\n", data)); err != nil {
					return fmt.Errorf("jsonl 파일 입력 오류 : %w\n", err)
				}
			}
		}
	}
	return nil
}

func (a *aiService) InjectInfoService(service types.RegionsService) {
	a.RegionsService = service
}
