package regions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
)

type regionsService struct {
	regionsRepository *RegionsRepository
}

func SetRegionsService(repository *RegionsRepository) types.RegionsService {
	r := &regionsService{
		regionsRepository: repository,
	}
	return r
}

/*
*

	흐름
	1. 지역, 카테고리 정보로 areaBasedList 호출 -> contentID 추출
	2. 추출 정보로 DetailCommon호출 -> 개요 추출, 여행 타이틀 추출, 카테고리 추출, 위치 추출(지역, 시군구 코드)
	3. 여행 타이틀, 위치(), 개요, 카테고리 추출

	서비스
	지역코드, 페이지네이션 혹은 개수 입력 받기
	-> 지역별로 여행지 정보 해당 수량대로 나옴
	-> 나오는대로 commonDetail API 호출해서 개요 추출, 여행 타이틀 추출, 카테고리 추출, 위치 추출(지역, 시군구 코드)
	-> 추출된 위치는

*
*/
func (r *regionsService) GetareaBasedList(areaCode types.AreaCode) ([]*types.AreaBasedListRes, error) {

	var result *types.DefaultTourAPIRes[*types.AreaBasedListRes]

	apiKey := os.Getenv("TOUR_API_KEY")
	if apiKey == "" {
		log.Fatal("환경 변수 'TOUR_API_KEY'가 설정되지 않았습니다.")
	}
	log.Println(apiKey)
	// Query 파라미터 설정
	v := url.Values{}                   // 인증키
	v.Set("numOfRows", "10")            // 데이터 개수
	v.Set("pageNo", "1")                // 페이지 번호
	v.Set("MobileOS", "ETC")            // OS 종류
	v.Set("MobileApp", "TestApp")       // 앱 이름
	v.Set("arrange", "D")               // 정렬 기준
	v.Set("areaCode", string(areaCode)) // 지역 코드
	v.Set("sigunguCode", "")            // 시군구 코드 (선택적)
	v.Set("cat1", "")                   // 대분류 (선택적)
	v.Set("cat2", "")                   // 중분류 (선택적)
	v.Set("cat3", "")
	v.Set("listYN", "Y")
	v.Set("_type", "json") // 소분류 (선택적)

	// 최종 URL 생성
	finalURL := fmt.Sprintf("%s?%s&serviceKey=%s", types.BASE_URL_AREA, v.Encode(), apiKey)
	log.Println("Requesting URL:", finalURL)

	res, err := http.Get(finalURL)
	if err != nil {
		return nil, fmt.Errorf("Tour API 요청 실패 : %w", err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Body 읽기 실패 : %w", err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("언마샬 실패 : %w", err)
	}

	return result.Response.Body.Items.Item, nil
}
