ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: dev-api
dev-api:
	go tool air -c .air.toml

.PHONY: run-api
run-api:
	go run ./cmd/api

.PHONY: test
test:
	go tool gotestsum -- ./...

.PHONY: test-verbose
test-verbose:
	go tool gotestsum --format testname -- ./... -v

.PHONY: test-cover
test-cover:
	go tool gotestsum -- ./... -cover

.PHONY: test-cover-profile
test-cover-profile:
	go test ./... -coverprofile=coverage.out

.PHONY: test-cover-html
test-cover-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: test-cover-func
test-cover-func:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy