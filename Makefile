.PHONY: run demo benchmark test clean

# Chạy demo chính
run:
	go run main.go

# Chạy demo với benchmark
demo:
	python3 scripts/run_demo.py

# Chỉ chạy benchmark
benchmark:
	go run cmd/benchmark/main.go

# Chạy tests
test:
	go test -v ./...

# Format code
fmt:
	go fmt ./...

# Clean build artifacts
clean:
	go clean
	go mod tidy

# Build binary
build:
	go build -o bin/patterns main.go

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run with race detection
race:
	go run -race main.go

# Generate documentation
docs:
	godoc -http=:6060

help:
	@echo "Available commands:"
	@echo "  run       - Chạy demo chính"
	@echo "  demo      - Chạy demo với Python script"
	@echo "  benchmark - Chạy performance benchmark"
	@echo "  test      - Chạy tests"
	@echo "  fmt       - Format code"
	@echo "  clean     - Clean artifacts"
	@echo "  build     - Build binary"
	@echo "  deps      - Install dependencies"
	@echo "  race      - Run with race detection"
	@echo "  docs      - Generate documentation" 