BINARY_NAME=jetea
SRC=main.go
BUILD_DIR=build
GO=go

all: build

build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC)

clean:
	rm -rf $(BUILD_DIR)

run: build
	$(BUILD_DIR)/$(BINARY_NAME)

fmt:
	$(GO) fmt ./...

lint:
	golangci-lint run

tidy: fmt lint mod-tidy

mod-tidy:
	$(GO) mod tidy

help:
	@echo "Makefile commands:"
	@echo "  make        - Build the binary"
	@echo "  make build  - Build the binary"
	@echo "  make clean  - Clean the build directory"
	@echo "  make run    - Build and run the binary"
	@echo "  make fmt    - Format Go code"
	@echo "  make lint   - Lint Go code"
	@echo "  make tidy   - Format, lint, and tidy Go modules"
	@echo "  make mod-tidy - Tidy Go modules"

.PHONY: all build clean run fmt lint tidy mod-tidy help
