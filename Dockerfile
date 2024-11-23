# 빌드 단계 (deploy-builder)
FROM golang:1.23.2-bullseye AS deploy-builder

WORKDIR /app

# 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o coursePickm

# 배포 이미지 (deploy)
FROM debian:bullseye-slim AS deploy

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# 빌드된 파일 복사
COPY --from=deploy-builder /app/coursePick ./

# 애플리케이션 실행
CMD ["./coursePick"]

# 개발용 이미지 (Dev)
FROM golang:1.23.2 AS dev

# 작업 디렉토리 설정
WORKDIR /app

# Swagger와 air 설치
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.11
RUN go install github.com/air-verse/air@latest

# 소스 코드 복사
COPY . .

# Swagger 문서 생성 후 air 실행
CMD ["sh", "-c", "swag init --parseDependency --parseInternal && air"]


