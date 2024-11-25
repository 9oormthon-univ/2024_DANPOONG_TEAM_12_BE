// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/courses/ai": {
            "post": {
                "description": "코스 추천이 성공적으로 완료되었습니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "코스 추천",
                "parameters": [
                    {
                        "description": "추천 요청 데이터",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RecommendCourseReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "성공 응답 데이터",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.Course"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "에러 응답 데이터",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Course": {
            "type": "object",
            "properties": {
                "end_date": {
                    "description": "종료 날짜, 시간 포인터 타입",
                    "type": "string"
                },
                "name": {
                    "description": "코스 이름, 최대 100자",
                    "type": "string"
                },
                "plans": {
                    "description": "연결된 계획들, 코스 삭제 시 함께 삭제",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Plan"
                    }
                },
                "start_date": {
                    "description": "시작 날짜, 시간 포인터 타입",
                    "type": "string"
                },
                "total_time": {
                    "description": "총 시간, 포인터 타입",
                    "type": "string"
                }
            }
        },
        "types.Place": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "장소 주소",
                    "type": "string"
                },
                "description": {
                    "description": "Duration  *int       // 소요 시간, 포인터 타입",
                    "type": "string"
                },
                "end_time": {
                    "description": "종료 시간",
                    "type": "string"
                },
                "image_url": {
                    "description": "이미지 URL, 최대 255자",
                    "type": "string"
                },
                "mapx": {
                    "description": "X 좌표 (경도 값, 위치 정보)",
                    "type": "string"
                },
                "mapy": {
                    "description": "Y 좌표 (위도 값, 위치 정보)",
                    "type": "string"
                },
                "start_time": {
                    "description": "시작 시간",
                    "type": "string"
                },
                "title": {
                    "description": "장소 이름, 최대 100자",
                    "type": "string"
                },
                "type": {
                    "description": "장소 유형",
                    "type": "string"
                }
            }
        },
        "types.Plan": {
            "type": "object",
            "properties": {
                "day_number": {
                    "description": "일자 번호",
                    "type": "string"
                },
                "places": {
                    "description": "연결된 장소들, 계획 삭제 시 함께 삭제",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Place"
                    }
                }
            }
        },
        "types.RecommendCourseReq": {
            "type": "object",
            "properties": {
                "area_code": {
                    "description": "지역 코드",
                    "type": "string"
                },
                "categories": {
                    "description": "카테고리",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "end_date": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "start_date": {
                    "description": "시작 날짜",
                    "type": "string"
                },
                "start_time": {
                    "description": "시작 시간",
                    "type": "string"
                },
                "total_time": {
                    "type": "string"
                },
                "with": {
                    "description": "누구와",
                    "type": "string"
                }
            }
        },
        "util.ResponseDTO": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Course Pick API",
	Description:      "지역 축제 및 여행 추천 서비스",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
