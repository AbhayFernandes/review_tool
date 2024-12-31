# Variables
GO_CMD = go
DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = docker-compose.yml
WEB_DIR = web
CLI_BINARY = cmd/crev/crev
API_DIR = ./cmd/api
CLI_DIR = ./cmd/crev
BUILD_DIR = ./build
JOB_PROCESSOR_DIR = ./cmd/job-processor
PKG_DIR = pkg
TEST_DIR = tests

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make build-cli           - Build the CLI binary"
	@echo "  make run-cli ARGS=...    - Run the CLI with arguments"
	@echo "  make run-api             - Run the API service"
	@echo "  make run-job-processor   - Run the Job Processor service"
	@echo "  make run-web             - Run the React frontend"
	@echo "  make build-docker        - Build all Docker images"
	@echo "  make build-web           - Build the web ui and run it"
	@echo "  make up                  - Start all services with Docker Compose"
	@echo "  make down                - Stop all services with Docker Compose"
	@echo "  make test                - Run all tests"
	@echo "  make clean               - Clean up generated files"
	@echo "  make build-cli-release   - Build the CLI binary for release"

# Build commands
.PHONY: build
build: build-api build-job-processor build-cli build-web
	@echo "All components built successfully!"

.PHONY: build-go
build-go: build-api build-job-processor build-cli
	@echo "All components built successfully!"

.PHONY: build-cli
build-cli:
	$(GO_CMD) build -o $(CLI_BINARY) $(CLI_DIR)

.PHONY: install-cli
install-cli:
	$(GO_CMD) install $(CLI_DIR)

.PHONY: build-api
build-api:
	$(GO_CMD) build -o $(BUILD_DIR)/api $(API_DIR)

.PHONY: build-job-processor
build-job-processor:
	$(GO_CMD) build -o $(BUILD_DIR)/api $(JOB_PROCESSOR_DIR)

.PHONY: build-web
build-web:
	cd $(WEB_DIR) && npm install && npm run build

.PHONY: build-docker
build-docker:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) build

.PHONY: build-cli-release
build-cli-release:
	$(GO_CMD) build -o $(BUILD_DIR)/crev $(CLI_DIR)

# Run commands
.PHONY: run-cli
run-cli: build-cli
	./$(CLI_BINARY) $(ARGS)

.PHONY: run-api
run-api:
	$(GO_CMD) run $(API_DIR)/main.go

.PHONY: run-job-processor
run-job-processor:
	$(GO_CMD) run $(JOB_PROCESSOR_DIR)/main.go

.PHONY: run-web
run-web:
	cd $(WEB_DIR) && npm run dev

.PHONY: up
up:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up --build -d

.PHONY: down
down:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

# Test commands
.PHONY: test
test:
	$(GO_CMD) test ./$(PKG_DIR)/ssh/... ./$(API_DIR)/... ./$(JOB_PROCESSOR_DIR)/... ./$(CLI_DIR)/...

.PHONY: test-cov
test-cov:
	@go list -f '{{.Dir}}' -m | \
	xargs -n1 | \
	sed 's|^/Users/[^/]*/workplace/review_tool/||' | \
	while read pkg; do \
		go test -coverprofile=$$(echo $$pkg | tr / -).cover /Users/$$USER/workplace/review_tool/$$pkg; \
	done
	@echo "mode: set" > c.out
	@grep -h -v "^mode:" ./*.cover >> c.out
	# @rm -f -- *.cover

# Generate protobuf files
.PHONY: proto
proto:
	cd "pkg/proto" && protoc echo.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

# Clean commands
.PHONY: clean
clean:
	rm -f $(CLI_BINARY)
	rm -rf $(BUILD_DIR)
	cd $(WEB_DIR) && rm -rf node_modules build
	$(GO_CMD) clean -cache -testcache
