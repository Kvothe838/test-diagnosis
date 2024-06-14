.PHONY: run
run: ## It runs the main app
	go run cmd/app/main.go

.PHONY: lint
lint: ## It starts the linter report
	@golangci-lint run --color always ./...


.PHONY: test
test: ## It runs the tests
	go test ./...