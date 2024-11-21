package regions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

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

// contentTypeID는 옵셔널로
func (r *regionsService) GetareaBasedList(areaCode types.AreaCode, contentTypeId types.ContentType) ([]*types.AreaBasedListRes, error) {
	v := url.Values{}
	if contentTypeId == "" {
		v.Set("contentTypeId", "")
	} else {
		v.Set("contentTypeId", string(contentTypeId))
	}

	var result *types.DefaultTourAPIRes[*types.AreaBasedListRes]

	apiKey := os.Getenv("TOUR_API_KEY")
	if apiKey == "" {
		log.Fatal("환경 변수 'TOUR_API_KEY'가 설정되지 않았습니다.")
	}
	log.Println(apiKey)
	// Query 파라미터 설정
	v.Set("numOfRows", "20")            // 데이터 개수
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

func (r *regionsService) GetDetailCommon(contentID string) (*types.DetailCommonRes, error) {

	var result *types.DefaultTourAPIRes[*types.DetailCommonRes]

	apiKey := os.Getenv("TOUR_API_KEY")
	if apiKey == "" {
		log.Fatal("환경 변수 'TOUR_API_KEY'가 설정되지 않았습니다.")
	}
	log.Println(apiKey)
	// Query 파라미터 설정
	v := url.Values{}
	v.Set("_type", "json")
	v.Set("MobileOS", "ETC") // OS 종류
	v.Set("MobileApp", "TestApp")
	v.Set("contentId", contentID)
	v.Set("overviewYN", "Y")
	v.Set("catcodeYN", "Y")
	v.Set("areacodeYN", "Y")

	// 최종 URL 생성
	finalURL := fmt.Sprintf("%s?%s&serviceKey=%s", types.BASE_URL_DETAIL_COMMON, v.Encode(), apiKey)
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

	// overview에서 \n, \t 제거
	if result != nil && len(result.Response.Body.Items.Item) > 0 {
		item := result.Response.Body.Items.Item[0]
		item.Overview = strings.ReplaceAll(item.Overview, "\n", "")
		item.Overview = strings.ReplaceAll(item.Overview, "\t", "")
	}

	// 인덱스가 원래 하나밖에 없음 사실상 슬라이스 x
	return result.Response.Body.Items.Item[0], nil
}

func (r *regionsService) GetAreaNameByCode(areaCode types.AreaCode) string {
	table := map[types.AreaCode]string{
		types.Seoul:     "서울",      // 서울특별시
		types.Incheon:   "인천",      // 인천광역시
		types.Daejeon:   "대전",      // 대전광역시
		types.Daegu:     "대구",      // 대구광역시
		types.Gwangju:   "광주",      // 광주광역시
		types.Busan:     "부산",      // 부산광역시
		types.Ulsan:     "울산",      // 울산광역시
		types.Sejong:    "세종",      // 세종특별자치시
		types.Gyeonggi:  "경기도",     // 경기도
		types.Gangwon:   "강원도",     // 강원도
		types.Chungbuk:  "충청북도",    // 충청북도
		types.Chungnam:  "충청남도",    // 충청남도
		types.Jeonbuk:   "전라북도",    // 전라북도
		types.Jeonnam:   "전라남도",    // 전라남도
		types.Gyeongbuk: "경상북도",    // 경상북도
		types.Gyeongnam: "경상남도",    // 경상남도
		types.Jeju:      "제주특별자치도", // 제주특별자치도
	}

	if location, found := table[areaCode]; found {
		return location
	}
	log.Fatalf("유효하지 않은 areaCode -> input : %s\n", areaCode)
	return ""
}

func (r *regionsService) GetContentTypeNameByCode(contentType types.ContentType) string {
	contentTypeMap := map[types.ContentType]string{
		types.ContentTypeTourism:       "관광지",
		types.ContentTypeCulture:       "문화시설",
		types.ContentTypeFestival:      "축제/공연/행사",
		types.ContentTypeTravelCourse:  "여행코스",
		types.ContentTypeLeisure:       "레포츠",
		types.ContentTypeAccommodation: "숙박",
		types.ContentTypeShopping:      "쇼핑",
		types.ContentTypeFood:          "음식",
	}
	if name, found := contentTypeMap[contentType]; found {
		return name
	}
	log.Fatalf("유효하지 않은 contentid -> input : %s\n", contentType)
	return ""
}
