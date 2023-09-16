.PHONY: fmt
fmt:
	@echo "Apply Format..."
	go fmt ./sample

.PHONY: fmt-check
fmt-check:
	@echo "Checking Format..."
	@if [ $(shell gofmt -s -l ./sample | wc -l) -gt 0 ]; then\
		exit 1;\
	fi

.PHONY: golangci-lint
golangci-lint:
	@echo "Checking Lint..."
	@ golangci-lint run ./sample/...

.PHONY: lint
lint: fmt-check golangci-lint
