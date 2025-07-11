APP_NAME=cash-log

.PHONY: all build run clean tidy help

all: build

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME)

run: build
	@echo "Running $(APP_NAME)..."
	./$(APP_NAME)

tidy:
	@echo "Tidying go.mod and go.sum..."
	go mod tidy

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

help:
	@echo "Makefile commands:"
	@echo "  make build   - Build the application"
	@echo "  make run     - Build and run the application"
	@echo "  make tidy    - Clean up go.mod and go.sum"
	@echo "  make clean   - Remove the built binary"
	@echo "  make help    - Show this help message"
