# .env 파일에서 환경 변수 불러오기
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DOCKER_TAG := latest

## 빌드 및 실행 관련 명령어
build: ## Build production Docker image
	docker build -t coursepick-backend:$(DOCKER_TAG) --target deploy ./

build-dev: ## Build development Docker image without cache
	docker compose build --no-cache

up: ## Start services with Docker Compose in detached mode
	docker compose up -d

down: ## Stop and remove all Docker containers
	docker compose down

restart: ## Restart Docker containers
	make down && make up

logs: ## Tail Docker Compose logs
	docker compose logs -f

ps: ## Check status of Docker containers
	docker compose ps

# migrate: ## Run database migrations
# 	mysqldef -u $(DB_USER) -p$(DB_PASSWORD) -h $(DB_HOST) -P $(DB_PORT) $(DB_NAME) < ./_tools/mysql/schema.sql
