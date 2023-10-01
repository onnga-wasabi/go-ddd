.PHONY: fmt
fmt:
	@echo "Applying format..."
	gofmt -w ./sample

.PHONY: fmt-check
fmt-check:
	@echo "Checking format..."
	@if [ $(shell gofmt -s -l ./sample | wc -l) -gt 0 ]; then\
		exit 1;\
	fi

.PHONY: golangci-lint
golangci-lint:
	@echo "Checking lint..."
	@golangci-lint run ./sample/...

.PHONY: lint
lint: fmt-check golangci-lint


.PHONY: create-migration
MIGRATE_NAME=""
create-migration:
	migrate create -ext sql -dir sample/infrastructure/migrations -seq ${MIGRATE_NAME}

.PHONY: migrate-up 
migrate-up:
	migrate -path sample/infrastructure/migrations -database postgres://postgres:pass@127.0.0.1:2345/postgres?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	migrate -path sample/infrastructure/migrations -database postgres://postgres:pass@127.0.0.1:2345/postgres?sslmode=disable down 1

.PHONY: install
install:
	@echo "Installing commands..."
	@# refer to https://golangci-lint.run/usage/install/
	@bash -c "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.54.2 > /dev/null 2>&1"
	@go install github.com/vektra/mockery/v2@v2.33.3
	@go install github.com/cweill/gotests/gotests@v1.6.0


LIBS := sample
go.mod: $(LIBS:%=%/*.go)
	@go mod tidy

DB_URI := "postgres://postgres:pass@127.0.0.1:2345/postgres?sslmode=disable"
data_model:
	xo schema $(DB_URI) -o sample/infrastructure/xo/persistence/data_model

mocks: go.mod
	@echo "Generating mocks..."
	@mockery

TEST_SOURCE_FILES := $(filter-out %_test.go %_mock.go %_xo.go, $(filter %go, $(shell find . -type f)))
tests: mocks go.mod
	@echo "Generating tests..."
	@gotests -all -w -template_dir gotests_template $(TEST_SOURCE_FILES)

.PHONY: run-test
run-test:
	@echo "Running tests..."
	@go test ./...
