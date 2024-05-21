MAIN_PACKAGE = main.go
BUILD = build
.PHONY: all
all: run

.PHONY: build
build:
	@echo "Building the project..."
	@go build -o $(BUILD)/app $(MAIN_PACKAGE)

.PHONY: run
run: build
	@echo "Running the project..."
	@./$(BUILD)/app

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v $(shell go list ./... | grep -v /vendor/)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f app
	@go clean -testcache

.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo ""
	@echo "Usage:"
	@echo "  make          Build and run the project"
	@echo "  make build    Build the project"
	@echo "  make run      Run the project"
	@echo "  make test     Run tests"
	@echo "  make clean    Clean up built files and test cache"
	@echo "  make help     Show this help message"
