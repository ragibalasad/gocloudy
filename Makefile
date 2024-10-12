BINARY_NAME=gocloudy
BUILD_DIR=bin
SOURCE_DIR=cmd/$(BINARY_NAME)
MAIN_FILE=main.go

VERSION := $(shell git describe --tags --abbrev=0)
BUILD_FLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)

	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -trimpath -o ./$(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_amd64 ./$(SOURCE_DIR)/
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -trimpath -o ./$(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_windows_amd64.exe ./$(SOURCE_DIR)/

	@echo "Binary built at $(BUILD_DIR)/"

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned up."
