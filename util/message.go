package util

// ResponseDTO 구조체 정의
type ResponseDTO struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 성공 응답 생성 함수
func SuccessResponse(message string, data interface{}) ResponseDTO {
	return ResponseDTO{
		Type:    "success",
		Message: message,
		Data:    data,
	}
}

// 에러 응답 생성 함수
func ErrorResponse(message string) ResponseDTO {
	return ResponseDTO{
		Type:    "error",
		Message: message,
	}
}
