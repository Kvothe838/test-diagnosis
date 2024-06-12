.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: lint
lint: ## It starts the linter report
	@golangci-lint run --color always ./...