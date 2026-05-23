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
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy 
