help:: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

storage-up: ## Initialize db in Docker
	@cd deployments && docker-compose up -d

storage-down: ## Finalize db in Docker
	@cd deployments && docker-compose down

storage-down-volume: ## Finalize db in Docker and remove data
	@cd deployments && docker-compose down -v

api-up: ## Run API
	@make storage-up
	@go run main.go

copy-env: ## Copy .env.example to projects root
	@cp ./configs/.env.example ./.env

test: ## Go test
	go test ./...

test_coverage: ## Go test with coverage file
	go test ./... -coverprofile=coverage.out

get_dependencies: ## Go get dependencies
	go get -v -t -d ./...

build-image: ## Build go image
	docker build -f deployments/Dockerfile -t lucassimao/my-api . 

deploy:
	docker build -f deployments/Dockerfile -t lucassimao/my-api . && docker push lucassimao/my-api