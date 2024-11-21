package ai

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
	"github.com/sashabaranov/go-openai"
)

type aiService struct {
	aiRepository *AIRepository
	types.RegionsService
}

func SetAIService(repository *AIRepository) types.AIService {
	r := &aiService{
		aiRepository: repository,
	}
	return r
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

	var chatFormatList []*types.ChatFormat

	for _, areaCode := range types.AreaCodes {
		log.Println("데이터 생성 중.....")
		areaList, err := a.RegionsService.GetareaBasedList(areaCode, "")
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
						Content: fmt.Sprintf("추천 장소: %s. 간단한 소개: %s", detail.Title, detail.Overview),
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
	return nil
}

func (a *aiService) InjectInfoService(service types.RegionsService) {
	a.RegionsService = service
}
