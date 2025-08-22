.PHONY: run
run:
	@go run ./cmd/pocketbook

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run -c ./.golangci.yml ./...

.PHONY: test
test:
	@go test ./... -v --cover
