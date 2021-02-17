DOCKER_COMPOSE_DEV=docker-compose -f docker-compose.yml -f docker-compose.dev.yml
DOCKER_COMPOSE_CI=docker-compose -f docker-compose.yml
.PHONY: test

help: ## Display available commands
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install and build go container
	$(DOCKER_COMPOSE_DEV) build

start: ## Start go script
	$(DOCKER_COMPOSE_DEV) up --force-recreate 

stop: ## Stop script
	$(DOCKER_COMPOSE_DEV) stop

test: ## Test the code
	$(DOCKER_COMPOSE_CI) run --rm --no-deps go go test ./src/...

test-verbose: ## Test the code
	$(DOCKER_COMPOSE_CI) run --rm --no-deps go go test -cover -v ./src/...
