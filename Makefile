GO_CMD = go
SERVICES = auth-service product-service checkout-service
BIN_DIR = bin

.PHONY: test build build-all run clean docker-build

# Run all tests
test:
	$(GO_CMD) test -v ./...

# Build a single service (usage: make build service=auth-service)
build:
	@test -n "$(service)" || (echo "Usage: make build service=auth-service" && exit 1)
	@mkdir -p $(BIN_DIR)
	$(GO_CMD) build -o $(BIN_DIR)/$(service) ./$(service)/cmd/main.go

# Build all services
build-all:
	@for s in $(SERVICES); do \
		$(MAKE) build service=$$s || exit 1; \
	done

# Run a service locally (usage: make run service=auth-service)
run:
	@test -n "$(service)" || (echo "Usage: make run service=auth-service" && exit 1)
	$(GO_CMD) run ./$(service)/cmd/main.go

# Clean build artifacts
clean:
	rm -rf $(BIN_DIR)

# Build Docker image (usage: make docker-build service=auth-service tag=1.0.0 IMAGE=username)
docker-build:
	@test -n "$(service)" || (echo "Usage: make docker-build service=auth-service tag=1.0.0 IMAGE=username" && exit 1)
	@test -n "$(tag)" || (echo "Usage: make docker-build service=auth-service tag=1.0.0 IMAGE=username" && exit 1)
	@test -n "$(IMAGE)" || (echo "Usage: make docker-build service=auth-service tag=1.0.0 IMAGE=username" && exit 1)
	docker build --build-arg SERVICE=$(service) -t $(IMAGE)/$(service):$(tag) .