PROJECT_PATH := "./cmd"

all: build

mod: ## run go mod
	GO111MODULE=on go mod tidy

build: ## Build the binary file
	@go install -v ${PROJECT_PATH}

test: ## Run unit tests
	@go test -short ./... -p 1

race: ## Run data race detector
	@go test -race -short ./... -p 1

build-docker-image: ## Build docker image
	@docker build \
		-t arithmetic:1.0.0 \
		-f dockerfile.arithmetic \
		.
	@docker system prune -f

up: ## start docker containers
	docker-compose up -d

down: ## stop docker containers
	docker-compose stop && docker-compose rm -f && docker system prune -f

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'