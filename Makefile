# Variables
BINARY_NAME=gocloudy
BUILD_DIR=bin
SOURCE_DIR=cmd/$(BINARY_NAME)
MAIN_FILE=main.go

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ./$(BUILD_DIR)/$(BINARY_NAME) ./$(SOURCE_DIR)/
	@echo "Binary built at $(BUILD_DIR)/$(BINARY_NAME)"

# Clean up build files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned up."

# Run the application
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Install the binary globally
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME) globally..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "$(BINARY_NAME) installed globally."

# Uninstall the binary
.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm /usr/local/bin/$(BINARY_NAME)
	@echo "$(BINARY_NAME) uninstalled."
