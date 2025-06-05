tidy:
	go mod tidy
	go fmt ./...
	fieldalignment -fix ./...
	go vet ./...
	golangci-lint run --fix ./...

run:
	make tidy
	go run main.go

build:
	make tidy
	go build

install_deps:
	# These needs sudo
	# apt install build-essential -y
    # curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install github.com/google/wire/cmd/wire@latest
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite

.PHONY: proto proto-clean proto-gen build run deps

# Configuration
PROTO_SRC_DIR := api/proto/src
PROTO_GEN_DIR := api/proto/gen
MICROSERVICES := $(notdir $(wildcard $(PROTO_SRC_DIR)/*))

# Proto generation
proto-clean:
	@echo "Cleaning generated proto files..."
	rm -rf $(PROTO_GEN_DIR)/*

proto-gen:
	@echo "Generating proto files..."
	@for service in $(MICROSERVICES); do \
		echo "Processing $$service..."; \
		mkdir -p $(PROTO_GEN_DIR)/$$service; \
		cd $(PROTO_SRC_DIR)/$$service && \
		protoc \
			--go_out=../../gen/$$service \
			--go_opt=paths=source_relative \
			--go-grpc_out=../../gen/$$service \
			--go-grpc_opt=paths=source_relative \
			*.proto; \
		cd - > /dev/null; \
	done

proto: proto-clean proto-gen

# Build and run
build: proto
	@echo "Building application..."
	go build -o bin/gateway main.go

run: build
	@echo "Running application..."
	./bin/gateway

# Dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Help
help:
	@echo "Available targets:"
	@echo "  proto-clean  - Clean generated proto files"
	@echo "  proto-gen    - Generate proto files"
	@echo "  proto        - Clean and generate proto files"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  deps         - Install dependencies"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Microservices found: $(MICROSERVICES)"


