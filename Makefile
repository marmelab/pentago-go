.PHONY: test

help: ## Display available commands
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install and build go container
	docker-compose build

start: ## Start go script
	docker-compose up --force-recreate -d

stop: ## Stop script
	docker-compose stop

test: ## Test the code
	docker-compose run --rm --no-deps go go test ./src/...

test-verbose: ## Test the code
	docker-compose run --rm --no-deps go go test -cover -v ./src/...
