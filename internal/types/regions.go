package types

import "time"

type RegionsService interface {
	GetAreaBasedList(areaCode AreaCode, contentTypeId ContentType) ([]*AreaBasedListRes, error)
	GetDetailCommon(contentID string) (*DetailCommonRes, error)
	GetAreaNameByCode(areaCode AreaCode) string
	GetContentTypeNameByCode(contentType ContentType) string
	GetAreaCodeByName(areaName string) AreaCode
	GetContentTypeCodeByName(contentTypeName string) ContentType
}
type ContentsService interface {
	GetAllContents() ([]Content, error)
	GetContentById(contentId int) (*Content, error)
}

// LocalInfo 구조체는 지역 정보를 나타냅니다.
type LocalInfo struct {
	LocalInfoID  int64  `gorm:"primaryKey;autoIncrement"` // 지역 정보 ID, 자동 증가 및 기본 키
	Name         string `gorm:"size:100;not null"`        // 지역 이름, 최대 100자, 필수 입력
	LocalFeature string `gorm:"type:text"`                // 지역 특징, 텍스트 타입
	Specialty    string `gorm:"type:text"`                // 지역 특산물, 텍스트 타입
	Symbol       string `gorm:"type:text"`                // 지역 상징, 텍스트 타입
	Attraction   string `gorm:"type:text"`                // 지역 명소, 텍스트 타입
	Activity     string `gorm:"type:text"`                // 지역 활동, 텍스트 타입
}

// Content 구조체는 콘텐츠 정보를 나타냅니다.
type Content struct {
	ContentID   int64         `gorm:"primaryKey;autoIncrement"`                            // 콘텐츠 ID, 자동 증가 및 기본 키
	Name        string        `gorm:"size:100;not null"`                                   // 콘텐츠 이름, 최대 100자, 필수 입력
	Type        string        `gorm:"type:enum('festival','event','attraction');not null"` // 콘텐츠 유형, 필수 입력
	Details     string        `gorm:"type:text"`                                           // 콘텐츠 세부 정보, 텍스트 타입
	StartDate   time.Time     `gorm:"type:date;not null"`                                  // 시작 날짜, 필수 입력
	EndDate     time.Time     `gorm:"type:date;not null"`                                  // 종료 날짜, 필수 입력
	Location    string        `gorm:"size:100"`                                            // 위치, 최대 100자
	CreatedAt   time.Time     `gorm:"autoCreateTime"`                                      // 생성 시간, 자동 생성
	UpdatedAt   time.Time     `gorm:"autoUpdateTime"`                                      // 수정 시간, 자동 업데이트
	LocalInfoID int64         `gorm:"not null"`                                            // 연결된 지역 정보 ID, 필수 입력
	Likes       []ContentLike `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE"`    // 연결된 좋아요, 콘텐츠 삭제 시 함께 삭제
}

// ContentLike 구조체는 콘텐츠에 대한 좋아요 정보를 나타냅니다.
type ContentLike struct {
	LikeID    int64     `gorm:"primaryKey;autoIncrement"`                          // 좋아요 ID, 자동 증가 및 기본 키
	UserID    int64     `gorm:"not null"`                                          // 사용자 ID, 필수 입력
	ContentID int64     `gorm:"not null"`                                          // 콘텐츠 ID, 필수 입력
	LikedAt   time.Time `gorm:"autoCreateTime"`                                    // 좋아요 생성 시간, 자동 생성
	User      User      `gorm:"foreignKey:UserID"`                                 // 사용자와의 관계
	Content   Content   `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"` // 콘텐츠와의 관계, 콘텐츠 삭제 시 좋아요도 삭제
}

// Tour API Dto ------------------ ------------------ ------------------ ------------------ ------------------ ------------------
type ContentType string
type Cat1 string
type Cat2 string
type Cat3 string
type AreaCode string
type DefaultTourAPIRes[T any] struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
		} `json:"header"`
		Body struct {
			Items struct {
				Item []T `json:"item"`
			} `json:"items"`
			NumOfRows  int `json:"numOfRows"`
			PageNo     int `json:"pageNo"`
			TotalCount int `json:"totalCount"`
		} `json:"body"`
	} `json:"response"`
}

// http://apis.data.go.kr/B551011/KorService1/areaBasedList1/
type AreaBasedListRes struct {
	Addr1         string      `json:"addr1"`         // 주소 1 (기본 주소 정보)
	Addr2         string      `json:"addr2"`         // 주소 2 (추가적인 주소 정보, 예: 건물 이름 등)
	AreaCode      AreaCode    `json:"areacode"`      // 지역 코드 (예: 서울, 부산 등과 같은 광역시나 도를 나타내는 코드)
	BookTour      string      `json:"booktour"`      // 예약 여부 (예약 필요 시 관련 정보, 비어 있을 수도 있음)
	Cat1          Cat1        `json:"cat1"`          // 대분류 카테고리 코드 (예: 관광지, 문화시설 등)
	Cat2          Cat2        `json:"cat2"`          // 중분류 카테고리 코드 (예: 관광 명소 중 특정 카테고리)
	Cat3          Cat3        `json:"cat3"`          // 소분류 카테고리 코드 (중분류 아래의 세부 카테고리)
	ContentID     string      `json:"contentid"`     // 콘텐츠 ID (특정 항목의 고유 식별자)
	ContentTypeID ContentType `json:"contenttypeid"` // 콘텐츠 유형 ID (예: 12는 관광지, 15는 축제, 32는 숙박시설)
	CreatedTime   string      `json:"createdtime"`   // 생성 시간 (항목 등록 시간, 형식: YYYYMMDDHHMMSS)
	FirstImage    string      `json:"firstimage"`    // 대표 이미지 URL (첫 번째 이미지)
	FirstImage2   string      `json:"firstimage2"`   // 추가 대표 이미지 URL (두 번째 이미지)
	CpyrhtDivCd   string      `json:"cpyrhtDivCd"`   // 저작권 구분 코드 (이미지나 정보에 대한 저작권 유형)
	MapX          string      `json:"mapx"`          // X 좌표 (경도 값, 위치 정보)
	MapY          string      `json:"mapy"`          // Y 좌표 (위도 값, 위치 정보)
	MLevel        string      `json:"mlevel"`        // 지도 수준 (지도 표시 시 확대 수준)
	ModifiedTime  string      `json:"modifiedtime"`  // 수정 시간 (마지막 수정 시간, 형식: YYYYMMDDHHMMSS)
	SigunguCode   string      `json:"sigungucode"`   // 시군구 코드 (해당 지역의 시, 군, 구 코드)
	Tel           string      `json:"tel"`           // 전화번호 (관련 문의 전화번호)
	Title         string      `json:"title"`         // 제목 (관광지, 행사 등의 이름)
	ZipCode       string      `json:"zipcode"`       // 우편번호 (해당 장소의 우편번호)
}

// http://apis.data.go.kr/B551011/KorService1/detailCommon1
type DetailCommonRes struct {
	ContentID     string `json:"contentid"`
	ContentTypeID string `json:"contenttypeid"`
	Title         string `json:"title"`
	CreatedTime   string `json:"createdtime,omitempty"`
	ModifiedTime  string `json:"modifiedtime,omitempty"`
	Tel           string `json:"tel,omitempty"`
	TelName       string `json:"telname,omitempty"`
	Homepage      string `json:"homepage,omitempty"`
	BookTour      string `json:"booktour,omitempty"`
	FirstImage    string `json:"firstimage,omitempty"`
	FirstImage2   string `json:"firstimage2,omitempty"`
	CpyrhtDivCd   string `json:"cpyrhtDivCd,omitempty"`
	AreaCode      string `json:"areacode,omitempty"`
	SigunguCode   string `json:"sigungucode,omitempty"`
	Cat1          string `json:"cat1,omitempty"`
	Cat2          string `json:"cat2,omitempty"`
	Cat3          string `json:"cat3,omitempty"`
	Addr1         string `json:"addr1,omitempty"`
	Addr2         string `json:"addr2,omitempty"`
	ZipCode       string `json:"zipcode,omitempty"`
	MapX          string `json:"mapx,omitempty"`
	MapY          string `json:"mapy,omitempty"`
	MLevel        string `json:"mlevel,omitempty"`
	Overview      string `json:"overview,omitempty"`
}

const (
	BASE_URL_AREA          = "http://apis.data.go.kr/B551011/KorService1/areaBasedList1"
	BASE_URL_DETAIL_COMMON = "http://apis.data.go.kr/B551011/KorService1/detailCommon1" // 공통 정보 조회
)

const (
	Seoul     AreaCode = "1"  // 서울특별시
	Incheon   AreaCode = "2"  // 인천광역시
	Daejeon   AreaCode = "3"  // 대전광역시
	Daegu     AreaCode = "4"  // 대구광역시
	Gwangju   AreaCode = "5"  // 광주광역시
	Busan     AreaCode = "6"  // 부산광역시
	Ulsan     AreaCode = "7"  // 울산광역시
	Sejong    AreaCode = "8"  // 세종특별자치시
	Gyeonggi  AreaCode = "31" // 경기도
	Gangwon   AreaCode = "32" // 강원도
	Chungbuk  AreaCode = "33" // 충청북도
	Chungnam  AreaCode = "34" // 충청남도
	Jeonbuk   AreaCode = "35" // 전라북도
	Jeonnam   AreaCode = "36" // 전라남도
	Gyeongbuk AreaCode = "37" // 경상북도
	Gyeongnam AreaCode = "38" // 경상남도
	Jeju      AreaCode = "39" // 제주특별자치도
)

var AreaCodes = []AreaCode{
	Seoul, Incheon, Daejeon, Daegu, Gwangju, Busan, Ulsan, Sejong,
	Gyeonggi, Gangwon, Chungbuk, Chungnam, Jeonbuk, Jeonnam, Gyeongbuk, Gyeongnam, Jeju,
}

const (
	ContentTypeTourism       ContentType = "12" // 관광지
	ContentTypeCulture       ContentType = "14" // 문화시설
	ContentTypeFestival      ContentType = "15" // 축제/공연/행사
	ContentTypeTravelCourse  ContentType = "25" // 여행코스
	ContentTypeLeisure       ContentType = "28" // 레포츠
	ContentTypeAccommodation ContentType = "32" // 숙박
	ContentTypeShopping      ContentType = "38" // 쇼핑
	ContentTypeFood          ContentType = "39" // 음식
)

var ContentTypeNames = []string{
	"관광지",
	"문화시설",
	"축제/공연/행사",
	"레포츠",
	"숙박",
	"쇼핑",
	"음식",
}

var ContentTypeCodes = []ContentType{
	ContentTypeTourism,
	ContentTypeCulture,
	ContentTypeFestival,
	ContentTypeTravelCourse,
	ContentTypeLeisure,
	ContentTypeAccommodation,
	ContentTypeShopping,
	ContentTypeFood,
}

type AreaBasedListRequest struct {
	AreaCode    string `json:"area_code"`
	ContentType string `json:"content_type"`
}
