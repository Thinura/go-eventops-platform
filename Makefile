LOCAL_BIN := $(CURDIR)/bin
AIR := $(LOCAL_BIN)/air

.PHONY: install-tools
install-tools:
	@mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/air-verse/air@latest

.PHONY: dev-api
dev-api: install-tools
	$(AIR) -c .air.toml

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