.PHONY: build run test lint clean dev-frontend dev-backend migrate-up migrate-down migrate-create docker-build

# Go parameters
BINARY_NAME=server
GO=go
GOFLAGS=-v

# Build the Go backend
build:
	$(GO) build $(GOFLAGS) -o bin/$(BINARY_NAME) ./cmd/server

# Run the Go backend
run: build
	./bin/$(BINARY_NAME)

# Run Go tests
test:
	$(GO) test ./... -v

# Run Go linter (requires golangci-lint)
lint:
	golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf web/dist/

# Start Vue dev server
dev-frontend:
	cd web && npm run dev

# Start Go backend in dev mode
dev-backend:
	$(GO) run ./cmd/server

# Database migrations (requires golang-migrate CLI)
migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down 1

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Build Docker image
docker-build:
	docker build -t warzone-stats-tracker .
