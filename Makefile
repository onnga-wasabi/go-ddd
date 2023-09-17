.PHONY: fmt
fmt:
	@echo "Apply Format..."
	gofmt -w ./sample

.PHONY: fmt-check
fmt-check:
	@echo "Checking Format..."
	@if [ $(shell gofmt -s -l ./sample | wc -l) -gt 0 ]; then\
		exit 1;\
	fi

.PHONY: golangci-lint
golangci-lint:
	@echo "Checking Lint..."
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
	go install github.com/vektra/mockery/v2@v2.33.3
	go install github.com/cweill/gotests/gotests@v1.6.0

mocks:
	@echo "Creating Mocks..."
	@mockery

tests: mocks
	@echo "Generating Tests..."
	gotests -all -w -template_dir gotests_template **/*.go

.PHONY: run-test
run-test:
	go test ./sample/...
