help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: ## Run linter
	go mod tidy
	golangci-lint run

test: ## Run go tests
	go test ./... -count=1
