package util

import (
	"fmt"
	"runtime"
)

// DetailedError 구조체 정의
type DetailedError struct {
	Keyword string // 에러 유형 (e.g., "DatabaseError", "ValidationError")
	Path    string // 파일 경로 + 함수명 + 라인 번호
	Message string // 에러 메시지
}

// Error 인터페이스 구현
func (e DetailedError) Error() string {
	return fmt.Sprintf("[%s] %s - %s", e.Keyword, e.Path, e.Message)
}

// 에러 생성 함수
func NewDetailedError(keyword, message string) DetailedError {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return DetailedError{Keyword: keyword, Path: "unknown", Message: message}
	}

	function := runtime.FuncForPC(pc).Name()
	path := fmt.Sprintf("%s:%d (%s)", file, line, function)

	return DetailedError{
		Keyword: keyword,
		Path:    path,
		Message: message,
	}
}

// 예시
// err = util.NewDetailedError("main err", err.Error())
// fmt.Println(err.Error())